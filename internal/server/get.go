package server

import (
	"configapi/pb"
	"context"
	"log"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

const updateInUseQuery = `UPDATE config SET in_use = CASE WHEN config_id = $1 THEN TRUE ELSE FALSE END`

const selectConfigByIdQuery = `SELECT config_id,name,config,version,in_use,created_at
		FROM config JOIN service ON config.service_id = service.service_id
		WHERE config_id = $1`

func (s *Server) Get(ctx context.Context, in *pb.ConfigID) (*pb.ConfigResponse, error) {
	return s.getConf(ctx, selectConfigByIdQuery, in.GetValue())
}

const selectUsingConfQuery = `SELECT config_id,name,config,version,in_use,created_at
		FROM config JOIN service ON config.service_id = service.service_id
		WHERE name = $1 AND in_use = TRUE`

func (s *Server) GetUsingConf(ctx context.Context, in *pb.Service) (*pb.ConfigResponse, error) {
	return s.getConf(ctx, selectUsingConfQuery, in.GetService())
}

func (s *Server) getConf(ctx context.Context, query string, arg any) (*pb.ConfigResponse, error) {
	t := time.Time{}
	out := pb.ConfigResponse{Config: &pb.Config{}}

	if err := s.pool.QueryRow(ctx, query,
		arg).Scan(&out.Id, &out.Config.Service, &out.Config.Config,
		&out.Version, &out.InUse, &t); err != nil {
		log.Println(err)
		return nil, err
	}
	out.CreatedAt = timestamppb.New(t)
	return &out, nil
}

const selectAllServiceConfQuery = `SELECT config_id,name,config,version,in_use,created_at
		FROM config JOIN service ON config.service_id = service.service_id
		WHERE name = $1`

func (s *Server) GetAllServiceConf(in *pb.Service, stream pb.ConfigService_GetAllServiceConfServer) error {
	rows, err := s.pool.Query(context.Background(), selectAllServiceConfQuery, in.GetService())
	if err != nil {
		log.Println(err)
		return err
	}
	defer rows.Close()

	t := time.Time{}
	out := pb.ConfigResponse{Config: &pb.Config{}}
	for rows.Next() {
		if err := rows.Scan(&out.Id, &out.Config.Service, &out.Config.Config,
			&out.Version, &out.InUse, &t); err != nil {
			log.Println(err)
			return err
		}
		out.CreatedAt = timestamppb.New(t)
		if err := stream.Send(&out); err != nil {
			log.Println(err)
			return err
		}
	}
	return nil
}
