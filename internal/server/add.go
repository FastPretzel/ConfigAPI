package server

import (
	"configapi/pb"
	"context"
	"log"
)

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
