package main

import (
	"context"
	"fmt"
	"log"

	"github.com/smallnest/rpcx/server"
	"rpcx-example/proto" // Protocol Buffers のパッケージをインポート
)

// HelloService: RPC サーバーのサービス
type HelloService struct{}

// SayHello: リモートで呼び出されるメソッド
func (h *HelloService) SayHello(ctx context.Context, args *proto.HelloRequest, reply *proto.HelloResponse) error {
	reply.Message = fmt.Sprintf("Hello, %s!", args.Name)
	return nil
}

func main() {
	s := server.NewServer()
	err := s.RegisterName("HelloService", new(HelloService), "")
	if err != nil {
		return
	} // サービスを登録

	fmt.Println("RPCサーバーを起動します...")
	if err := s.Serve("tcp", ":8972"); err != nil {
		log.Fatalf("サーバー起動エラー: %v", err)
	}
}
