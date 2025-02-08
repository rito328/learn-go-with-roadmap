package main

import (
	"context"
	"fmt"
	"github.com/smallnest/rpcx/protocol"
	"log"

	"github.com/smallnest/rpcx/client"
	"rpcx-example/proto" // Protocol Buffers のパッケージをインポート
)

func main() {
	d, _ := client.NewPeer2PeerDiscovery("tcp@localhost:8972", "")
	client.DefaultOption.SerializeType = protocol.ProtoBuffer // シリアライズ方式を protobuf に設定
	xclient := client.NewXClient("HelloService", client.Failover, client.RandomSelect, d, client.DefaultOption)
	defer xclient.Close()

	args := &proto.HelloRequest{Name: "Go Developer"}
	reply := &proto.HelloResponse{}

	// RPC メソッド呼び出し
	err := xclient.Call(context.Background(), "SayHello", args, reply)
	if err != nil {
		log.Fatalf("RPCエラー: %v", err)
	}

	fmt.Println("サーバーからのレスポンス:", reply.Message)
}
