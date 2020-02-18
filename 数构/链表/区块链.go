package main

import (
	"crypto/sha256"
	"encoding/json"
	"flag"
	"fmt"
	"golang.org/x/net/websocket"
	"io"
	"log"
	"net/http"
	"sort"
	"strings"
	"time"
)

/**
* @Author : awen
* @Date : 2020/1/20 5:22 下午
 */

const (
	queryLatest = iota
	queryAll
	responseBlockchain
)

var (
	sockets      []*websocket.Conn
	blockchain   = []*Block{genesisBlock}
	initialPeers = flag.String("peers", "ws://localhost:6001", "initial peers")
	httpAddr     = flag.String("api", ":3001", "api server address.")
	p2pAddr      = flag.String("p2p", ":6001", "p2p server address.")
)

// 定义创世区块
var genesisBlock = &Block{
	Index:        0,
	PreviousHash: "0",
	Timestamp:    1465154705,
	Data:         "my genesis block!!",
	Hash:         "816534932c2b7154836da6afc367695e6337db8a921823784c14378abed4f7d7",
}

// 定义区块结构体
type Block struct {
	Index        int64  // 区块下标
	PreviousHash string // 前一个区块的哈希值
	Timestamp    int64  // 时间戳
	Data         string // 数据
	Hash         string // 当前区块哈希

}

// 格式化字符串输出
func (b *Block) String() string {
	return fmt.Sprintf("index: %d,previousHash:%s,timestamp:%d,data:%s,hash:%s", b.Index, b.PreviousHash, b.Timestamp, b.Data, b.Hash)
}

type ByIndex []*Block

// 实现数组方法
func (b ByIndex) Len() int {
	return len(b)
}
func (b ByIndex) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}
func (b ByIndex) Less(i, j int) bool {
	return b[i].Index < b[j].Index
}

// 响应的区块链
type ResponseBlockChain struct {
	Type int    `json:"type"` // 区块类型
	Data string `json:"data"`
}

func ErrFatal(msg string, err error) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}

/**
 * @Description: 基于 websocket 对每一个地址初始化连接
 * @Param: peerAddrs
 * @Return:
 * @Date: 2020/1/20
 * @Time: 5:37 下午
 */
func connectToPeers(peerAddrs []string) {
	for _, peer := range peerAddrs {
		if peer == "" {
			continue
		}
		ws, err := websocket.Dial(peer, "", peer)
		if err != nil {
			log.Println("dial to peer ", err)
			continue
		}
		initConnections(ws)
	}
}

/**
 * @Description: 初始化连接
 * @Param: ws-a client
 * @Return:
 * @Date: 2020/1/20
 * @Time: 5:43 下午
 */
func initConnections(ws *websocket.Conn) {
	// 此协程会一直跑
	go wsHandleP2P(ws)
	log.Println("query latest block.")
	ws.Write(queryLatestMsg())
}

/**
 * @Description: 处理p2p网络
 * @Param:
 * @Return:
 * @Date: 2020/1/20
 * @Time: 5:48 下午
 */
func wsHandleP2P(ws *websocket.Conn) {
	var (
		v    = &ResponseBlockChain{}
		peer = ws.LocalAddr().String()
	)
	// 	添加客户端
	sockets = append(sockets, ws)

	for {
		var msg []byte
		err := websocket.Message.Receive(ws, &msg) // 不断的接受二进制数据
		if err == io.EOF {
			log.Printf("p2p Peer[%s] shutdown, remove it form peers pool.\n", peer)
			break
		}
		if err != nil { // 接受出错了
			log.Println("Can't receive p2p msg from ", peer, err.Error())
			break
		}
		log.Printf("Received[from %s]: %s.\n", peer, msg)
		err = json.Unmarshal(msg, v) //信息映射到 v 反序列化
		ErrFatal("invalid p2p msg", err)

		switch v.Type {
		case queryLatest:
			v.Type = responseBlockchain

			bs := responseLatestMsg()
			log.Printf("responseLatestMsg: %s\n", bs)
			ws.Write(bs)
		case queryAll:
			d, _ := json.Marshal(blockchain)
			v.Type = responseBlockchain
			v.Data = string(d)
			bs, _ := json.Marshal(v)
			log.Printf("responseChainMsg: %s\n", bs)
			ws.Write(bs)
		case responseBlockchain:
			handleBlockchainResponse([]byte(v.Data))
		}
	}
}

/**
 * @Description: 处理区块链响应
 * @Param:
 * @Return:
 * @Date: 2020/1/20
 * @Time: 6:02 下午
 */
func handleBlockchainResponse(msg []byte) {
	var receivedBlocks = []*Block{}
	err := json.Unmarshal(msg, &receivedBlocks) // 反序列化
	ErrFatal("invalid blockchain", err)
	sort.Sort(ByIndex(receivedBlocks))
	latestBlockReceived := receivedBlocks[len(receivedBlocks)-1]
	latestBlockHeld := getLatestBlock()
	if latestBlockReceived.Index > latestBlockHeld.Index {
		log.Printf("blockchain possibly behind. We got: %d Peer got: %d", latestBlockHeld.Index, latestBlockReceived.Index)
		if latestBlockHeld.Hash == latestBlockReceived.PreviousHash {
			log.Println("We can append the received block to our chain.")
			blockchain = append(blockchain, latestBlockReceived)
		} else if len(receivedBlocks) == 1 {
			log.Println("We have to query the chain from our peer.")
			broadcast(queryAllMsg())
		} else {
			log.Println("Received blockchain is longer than current blockchain.")
			replaceChain(receivedBlocks)
		}
	} else {
		log.Println("received blockchain is not longer than current blockchain. Do nothing.")
	}
}

func replaceChain(blocks []*Block) {
	if isValidChain(blocks) && len(blocks) > len(blockchain) {
		log.Println("Received blockchain is valid. Replacing current blockchain with received blockchain.")
		blockchain = blocks
		broadcast(responseLatestMsg())
	} else {
		log.Println("Received blockchain invalid.")
	}
}

/**
 * @Description: 区块链是否合法
 * @Param:
 * @Return:
 * @Date: 2020/1/20
 * @Time: 6:29 下午
 */
func isValidChain(blocks []*Block) bool {
	if blocks[0].String() != genesisBlock.String() {
		log.Println("No same GenesisBlock.", blocks[0].String())
		return false
	}
	var temp = []*Block{blocks[0]}
	for i := 1; i < len(blocks); i++ {
		if isValidNewBlock(blocks[i], temp[i-1]) {
			temp = append(temp, blocks[i])
		} else {
			return false
		}
	}
	return true
}

/**
 * @Description: 新区块是否合法
 * @Param:
 * @Return:
 * @Date: 2020/1/20
 * @Time: 6:31 下午
 */
func isValidNewBlock(nb, pb *Block) (ok bool) {
	if nb.Hash == calculateHashForBlock(nb) &&
		pb.Index+1 == nb.Index &&
		pb.Hash == nb.PreviousHash {
		ok = true
	}
	return
}

/**
 * @Description: 计算当前区块哈希
 * @Param:
 * @Return:
 * @Date: 2020/1/20
 * @Time: 6:31 下午
 */
func calculateHashForBlock(b *Block) string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(fmt.Sprintf("%d%s%d%s", b.Index, b.PreviousHash, b.Timestamp, b.Data))))
}

/**
 * @Description: 向所有客户端广播📢
 * @Param:
 * @Return:
 * @Date: 2020/1/20
 * @Time: 6:26 下午
 */
func broadcast(msg []byte) {
	for n, socket := range sockets {
		_, err := socket.Write(msg)
		if err != nil {
			log.Printf("peer [%s] disconnected.", socket.RemoteAddr().String())
			sockets = append(sockets[0:n], sockets[n+1:]...)
		}
	}
}

// 查询所有信息
func queryAllMsg() []byte {
	return []byte(fmt.Sprintf("{\"type\": %d}", queryAll))
}

/**
 * @Description: 获取最新区块
 * @Param:
 * @Return:
 * @Date: 2020/1/20
 * @Time: 6:23 下午
 */
func getLatestBlock() (block *Block) {
	return blockchain[len(blockchain)-1]
}

/**
 * @Description: 响应最新的信息
 * @Param:
 * @Return:
 * @Date: 2020/1/20
 * @Time: 5:57 下午
 */
func responseLatestMsg() (bs []byte) {
	var v = &ResponseBlockChain{
		Type: responseBlockchain,
	}
	d, _ := json.Marshal(blockchain[len(blockchain)-1:])
	v.Data = string(d)
	bs, _ = json.Marshal(v) // 结构体 --> 字节流  序列化
	return
}

/**
 * @Description: 查询区块最新信息
 * @Param:
 * @Return:
 * @Date: 2020/1/20
 * @Time: 5:44 下午
 */
func queryLatestMsg() []byte {
	return []byte(fmt.Sprintf("{\"type\": %d}", queryLatest))
}

/**
 * @Description: 生成下一个区块
 * @Param:
 * @Return:
 * @Date: 2020/1/20
 * @Time: 6:32 下午
 */
func generateNextBlock(data string) (nb *Block) {
	var previousBlock = getLatestBlock()
	nb = &Block{
		Data:         data,
		PreviousHash: previousBlock.Hash,
		Index:        previousBlock.Index + 1,
		Timestamp:    time.Now().Unix(),
	}
	nb.Hash = calculateHashForBlock(nb)
	return
} /**
 * @Description: 新增区块
 * @Param:
 * @Return: nil
 * @Date: 2020/1/20
 * @Time: 6:32 下午
 */
func addBlock(b *Block) {
	if isValidNewBlock(b, getLatestBlock()) {
		blockchain = append(blockchain, b)
	}
}

/**
接口方法
*/
func handleBlocks(w http.ResponseWriter, r *http.Request) {
	bs, _ := json.Marshal(blockchain)
	w.Write(bs)
}
func handleMineBlock(w http.ResponseWriter, r *http.Request) {
	var v struct {
		Data string `json:"data"`
	}
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()
	err := decoder.Decode(&v)
	if err != nil {
		w.WriteHeader(http.StatusGone)
		log.Println("[API] invalid block data : ", err.Error())
		w.Write([]byte("invalid block data. " + err.Error() + "\n"))
		return
	}
	block := generateNextBlock(v.Data)
	addBlock(block)
	broadcast(responseLatestMsg())
}
func handlePeers(w http.ResponseWriter, r *http.Request) {
	var slice []string
	for _, socket := range sockets {
		if socket.IsClientConn() {
			slice = append(slice, strings.Replace(socket.LocalAddr().String(), "ws://", "", 1))
		} else {
			slice = append(slice, socket.Request().RemoteAddr)
		}
	}
	bs, _ := json.Marshal(slice)
	w.Write(bs)
}
func handleAddPeer(w http.ResponseWriter, r *http.Request) {
	var v struct {
		Peer string `json:"peer"`
	}
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()
	err := decoder.Decode(&v)
	if err != nil {
		w.WriteHeader(http.StatusGone)
		log.Println("[API] invalid peer data : ", err.Error())
		w.Write([]byte("invalid peer data. " + err.Error()))
		return
	}
	connectToPeers([]string{v.Peer})
}

func main() {
	flag.Parse() //解析命令行参数
	connectToPeers(strings.Split(*initialPeers, ","))
	http.HandleFunc("/blocks", handleBlocks)
	http.HandleFunc("/mine_block", handleMineBlock)
	http.HandleFunc("/peers", handlePeers)
	http.HandleFunc("/add_peer", handleAddPeer)
	go func() {
		log.Println("Listen HTTP on", *httpAddr) // http服务
		ErrFatal("start api server", http.ListenAndServe(*httpAddr, nil))
	}()
	http.Handle("/", websocket.Handler(wsHandleP2P))
	log.Println("Listen P2P on ", *p2pAddr)
	ErrFatal("start p2p server", http.ListenAndServe(*p2pAddr, nil)) // p2p服务
}
