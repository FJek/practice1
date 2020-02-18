package main

import (
	"context"
	"github.com/smallnest/rpcx/client"
	"log"
)

/**
* @Author : awen
* @Date : 2020/2/5 8:57 下午
 */

func main() {
	// 定义了使用什么方式来实现服务发现。 在这里我们使用最简单的 Peer2PeerDiscovery（点对点）。客户端直连服务器来获取服务地址。
	d := client.NewPeer2PeerDiscovery("tcp@localhost", "")
	//  创建了 XClient， 并且传进去了 FailMode、 SelectMode 和默认选项。
	//  FailMode 告诉客户端如何处理调用失败：重试、快速返回，或者 尝试另一台服务器。
	//  SelectMode 告诉客户端如何在有多台服务器提供了同一服务的情况下选择服务器。
	xclient := client.NewXClient("Arith", client.Failtry, client.RandomSelect, d, client.DefaultOption)
	defer xclient.Close()

	// 定义了请求：这里我们想获得 10 * 20 的结果。 当然我们可以自己算出结果是 200，但是我们仍然想确认这与服务器的返回结果是否一致。
	args := &Args{
		A: 10,
		B: 20,
	}

	// 定义了响应对象， 默认值是0值， 事实上 rpcx 会通过它来知晓返回结果的类型，然后把结果反序列化到这个对象
	reply := &Reply{}

	//  调用了远程服务并且同步获取结果。
	err := xclient.Call(context.Background(), "Mul", args, reply)
	if err != nil {
		log.Fatalf("failed to call: %v", err)
	}

	log.Printf("%d * %d = %d", args.A, args.B, reply.C)
}

// 异步调用
func syncCall()  {
	d := client.NewPeer2PeerDiscovery("tcp@", "")
	xclient := client.NewXClient("Arith", client.Failtry, client.RandomSelect, d, client.DefaultOption)
	defer xclient.Close()

	args := &Args{
		A: 10,
		B: 20,
	}

	reply := &Reply{}
	call, err := xclient.Go(context.Background(), "Mul", args, reply, nil)
	if err != nil {
		log.Fatalf("failed to call: %v", err)
	}

	replyCall := <-call.Done
	if replyCall.Error != nil {
		log.Fatalf("failed to call: %v", replyCall.Error)
	} else {
		log.Printf("%d * %d = %d", args.A, args.B, reply.C)
	}
}
