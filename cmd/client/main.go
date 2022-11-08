package main

import (
	"configapi/pb"
	"context"
	"fmt"
	"io"
	"log"

	"google.golang.org/grpc"
)

// FOR TESTING
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
	fmt.Printf("Add response: %v\n", idx.GetValue())

	log.Println("exec get")
	out, err := c.Get(context.Background(), &pb.ConfigID{Value: idx.GetValue()})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Get response: %v, %v, %v, %v, %v, %v\n", out.GetId(),
		out.GetConfig().Service, out.GetConfig().Config, out.GetVersion(), out.GetInUse(), out.GetCreatedAt().AsTime())

	log.Println("exec getAll")
	stream, err := c.GetAllServiceConf(context.Background(), &pb.Service{Service: "k8s"})
	if err != nil {
		log.Fatal(err)
	}
	for {
		out, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("%v.ListFeatures(_) = _, %v", c, err)
		}
		fmt.Printf("GetAll response: %v, %v, %v, %v, %v, %v\n", out.GetId(),
			out.GetConfig().Service, out.GetConfig().Config, out.GetVersion(), out.GetInUse(), out.GetCreatedAt().AsTime())
	}

	log.Println("exec getUsing")
	out, err = c.GetUsingConf(context.Background(), &pb.Service{Service: "k8s"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("GetUsing response: %v, %v, %v, %v, %v, %v\n", out.GetId(),
		out.GetConfig().Service, out.GetConfig().Config, out.GetVersion(), out.GetInUse(), out.GetCreatedAt().AsTime())

	log.Println("exec update")
	update := `{"key1": "VALUE1","key2": "VALUE2"}`
	out, err = c.Update(context.Background(), &pb.UpdateConfig{Id: 1, Config: &update})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Update response: %v, %v, %v, %v, %v, %v\n", out.GetId(),
		out.GetConfig().Service, out.GetConfig().Config, out.GetVersion(), out.GetInUse(), out.GetCreatedAt().AsTime())

	log.Println("exec delete")
	outDelete, err := c.DeleteConf(context.Background(), &pb.ConfigID{Value: 1})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("DeleteConf response: %v\n", outDelete.GetSuccess())

	log.Println("exec deleteService")
	outDelete, err = c.DeleteService(context.Background(), &pb.Service{Service: "k8s"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("DeleteService response: %v\n", outDelete.GetSuccess())
	fmt.Println("END")
}
