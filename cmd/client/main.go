package main

import (
	"configapi/pb"
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:5000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("can't connect to server: %v", err)
	}
	defer conn.Close()

	c := pb.NewConfigServiceClient(conn)
	log.Printf("created client: %v", c)

	log.Println("exec add")
	idx, err := c.Add(context.Background(), &pb.Config{Service: "k8s",
		Config: `{"key1": "value1","key2": "value2"}`})
	if err != nil {
		log.Fatal(err)
	}
	log.Println("exec get")
	out, err := c.Get(context.Background(), &pb.ConfigID{Value: idx.GetValue()})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Response: %v, %v, %v, %v, %v, %v\n", out.GetId(),
		out.GetConfig().Service, out.GetConfig().Config, out.GetVersion(), out.GetInUse(), out.GetCreatedAt().AsTime())
}
