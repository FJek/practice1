package main

import (
	"context"
	"github.com/smallnest/rpcx/server"
)

/**
* @Author : awen
* @Date : 2020/2/5 8:51 下午
 */

type Args struct {
	A int
	B int
}

type Reply struct {
	C int
}
type Arith int

// 乘法服务
func (t *Arith) Mul(ctx context.Context, args *Args, reply *Reply) error {
	reply.C = args.A * args.B
	return nil
}

/**
 * @Description:
 * @Param:
 * @Return:
 * @Date: 2020/2/5
 * @Time: 8:52 下午
 */
func main() {
	// 注册服务
	s := server.NewServer()
	s.RegisterName("Arith", new(Arith), "")
	s.Serve("tcp", ":8972")
}
