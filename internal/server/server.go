package server

import (
	"configapi/pb"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Server struct {
	pb.UnsafeConfigServiceServer
	pool *pgxpool.Pool
}

func NewServer(pgxPool *pgxpool.Pool) *Server {
	return &Server{pool: pgxPool}
}
