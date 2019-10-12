package main

import (
	"context"
	"fmt"
	"github.com/tchaudhry91/wts-service/svc/pb"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:7878", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	client := pb.NewWTSClient(conn)
	resp, err := client.Get(context.Background(), &pb.GetRequest{Name: "Hello", User: "Me"})
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v", resp.Record)
}
