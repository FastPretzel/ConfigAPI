package server

import (
	"configapi/pb"
	"context"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Server struct {
	pb.UnsafeConfigServiceServer
	pool *pgxpool.Pool
}

func NewServer(pgxPool *pgxpool.Pool) *Server {
	return &Server{pool: pgxPool}
}

const insertServiceQuery = `INSERT INTO service (name) VALUES ($1)
		ON CONFLICT (name) DO UPDATE SET name = EXCLUDED.name RETURNING service_id`

const insertConfigQuery = `INSERT INTO config (service_id,config,version,created_at)
		VALUES ($1,$2,(SELECT COUNT(*) FROM config WHERE service_id = $3)+1,
		CURRENT_TIMESTAMP(2)) RETURNING config_id`

func (s *Server) Add(ctx context.Context, in *pb.Config) (*pb.ConfigID, error) {
	var (
		serviceId int64
		confId    int64
	)
	if err := s.pool.QueryRow(ctx, insertServiceQuery,
		in.GetService()).Scan(&serviceId); err != nil {
		log.Println(err)
		return nil, err
	}
	if err := s.pool.QueryRow(ctx, insertConfigQuery, serviceId,
		in.GetConfig(), serviceId).Scan(&confId); err != nil {
		log.Println(err)
		return nil, err
	}
	return &pb.ConfigID{Value: confId}, nil
}

// func (s *Server) insertConfig(ctx context.Context,)
const updateInUseQuery = `UPDATE config SET in_use = CASE WHEN config_id = $1 THEN TRUE ELSE FALSE END`

const selectConfigByIdQuery = `SELECT config_id,name,config,version,in_use,created_at
		FROM config JOIN service ON config.service_id = service.service_id
		WHERE config_id = $1`

func (s *Server) Get(ctx context.Context, in *pb.ConfigID) (*pb.ConfigResponse, error) {
	t := time.Time{}
	out := pb.ConfigResponse{Config: &pb.Config{}}
	if err := s.pool.QueryRow(ctx, selectConfigByIdQuery,
		in.GetValue()).Scan(&out.Id, &out.Config.Service, &out.Config.Config,
		&out.Version, &out.InUse, &t); err != nil {
		log.Println(err)
		return nil, err
	}
	out.CreatedAt = timestamppb.New(t)
	if _, err := s.pool.Exec(ctx, updateInUseQuery,
		in.GetValue()); err != nil {
		return nil, err
	}
	return &out, nil
}

func (s *Server) DeleteConf(ctx context.Context, in *pb.ConfigID) (*pb.DeleteResponse, error) {
	_, err := s.pool.Exec(ctx, "DELETE FROM config WHERE config_id = $1", in.GetValue())
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &pb.DeleteResponse{Success: true}, nil
}

func (s *Server) Update(ctx context.Context, in *pb.UpdateConfig) (*pb.ConfigResponse, error) {
	rows, err := s.pool.Query(ctx, "UPDATE config SET config = $1 WHERE config_id = $2 RETURNING *",
		in.GetConfig(), in.GetId())
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()

	t := time.Time{}
	out := pb.ConfigResponse{Config: &pb.Config{}}
	for rows.Next() {
		if err := rows.Scan(&out.Id, &out.Config.Service,
			&out.Config.Config, &t); err != nil {
			log.Println(err)
			return nil, err
		}
	}
	out.CreatedAt = timestamppb.New(t)
	return &out, nil
}
