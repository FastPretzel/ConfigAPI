package main

import (
	"configapi/internal/client/domain"
	repo "configapi/internal/client/repository"
	"configapi/pb"
	"context"
	"flag"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("localhost:5000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("can't connect to server: %v", err)
	}
	defer conn.Close()

	c := pb.NewConfigServiceClient(conn)

	client := domain.Client{
		ClientProtobuf: c,
		Repo:           repo.NewRepository()}
	log.Printf("created client: %v", client)

	flag.Parse()
	if err := methodMux(context.Background(), &client); err != nil {
		log.Fatal(err)
	}
}

func methodMux(ctx context.Context, client *domain.Client) error {
	method := client.Repo.GetMethod()
	switch method {
	case "Add":
		return client.Add(context.Background())
	case "Get":
		return client.Get(context.Background())
	case "GetUsingConf":
		return client.GetUsingConf(context.Background())
	case "GetAllServiceConf":
		return client.GetAllServiceConf(context.Background())
	case "Use":
		return client.Use(context.Background())
	case "Update":
		return client.Update(context.Background())
	case "DeleteConf":
		return client.DeleteConf(context.Background())
	case "DeleteService":
		return client.DeleteService(context.Background())
	default:
		return fmt.Errorf("Invalid method: %v\n", method)
	}
}
