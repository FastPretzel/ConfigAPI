package server

import (
	"configapi/pb"
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Server struct {
	pb.UnsafeConfigServiceServer
	pool *pgxpool.Pool
}

func NewServer(pgxPool *pgxpool.Pool) *Server {
	return &Server{pool: pgxPool}
}

func (s *Server) Add(ctx context.Context, in *pb.Config) (*pb.ConfigID, error) {
	rows, err := s.pool.Query(ctx, "INSERT INTO config "+
		"(service,config,created_at)) VALUES ($1,$2,CURRENT_TIMESTAMP(2))"+
		" RETURNING config_id", in.GetService(), in.GetConfig())
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()
	var id int64
	for rows.Next() {
		if err := rows.Scan(&id); err != nil {
			log.Println(err)
			return nil, err
		}
	}
	return &pb.ConfigID{Value: id}, nil
}

func (s *Server) Get(ctx context.Context, in *pb.ConfigID) (*pb.ConfigResponse, error) {
	rows, err := s.pool.Query(ctx, "SELECT * FROM config WHERE config_id = $1",
		in.GetValue())
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()
	out := pb.ConfigResponse{}
	for rows.Next() {
		if err := rows.Scan(&out.Id, &out.Config, &out.CreatedAt); err != nil {
			log.Println(err)
			return nil, err
		}
	}
	return &out, nil
}

func (s *Server) Delete(ctx context.Context, in *pb.ConfigID) (*pb.DeleteResponse, error) {
	_, err := s.pool.Exec(ctx, "DELETE FROM config WHERE config_id = $1", in.GetValue())
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &pb.DeleteResponse{Success: true}, nil
}

func (s *Server) Update(ctx context.Context, in *pb.UpdateConfig) (*pb.ConfigResponse, error) {
	//...
	return &pb.ConfigResponse{}, nil
}
