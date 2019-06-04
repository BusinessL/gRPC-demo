package main

import (
	"context"
	"log"

	pb "gRPC-demo/echospec"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:9999", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("连线失败：%v", err)
	}
	defer conn.Close()

	c := pb.NewEchoClient(conn)

	r, err := c.Echo(context.Background(), &pb.EchoRequest{Msg: "bag"})
	if err != nil {
		log.Fatalf("无法执行：%v", err)
	}
	log.Printf("回传結果：%s , 时间:%d", r.Msg, r.Unixtime)

}