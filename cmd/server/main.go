package main

import (
	srv "configapi/internal/server"
	"configapi/pb"
	"configapi/pkg/postgres"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	url := "postgres://postgres:secret@db:5432/postgres"
	pool, err := postgres.NewPool(url)
	if err != nil {
		log.Fatal(err)
	}
	defer pool.Close()

	listen, err := net.Listen("tcp", ":5000")
	if err != nil {
		log.Fatal(err)
	}
	s := grpc.NewServer()

	pb.RegisterConfigServiceServer(s, srv.NewServer(pool))
	if err = s.Serve(listen); err != nil {
		log.Fatal(err)
	}
}
