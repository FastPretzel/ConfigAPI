package server

import (
	"configapi/pb"
	"context"
	"log"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

const updateConfQuery = `UPDATE config SET config = $1 FROM service
		WHERE service.service_id = config.service_id AND config_id = $2
		RETURNING config_id,name,config,version,in_use,created_at`

func (s *Server) Update(ctx context.Context, in *pb.UpdateConfig) (*pb.ConfigResponse, error) {
	t := time.Time{}
	out := pb.ConfigResponse{Config: &pb.Config{}}
	if err := s.pool.QueryRow(ctx, updateConfQuery,
		in.GetConfig(), in.GetId()).Scan(&out.Id, &out.Config.Service,
		&out.Config.Config, &out.Version, &out.InUse, &t); err != nil {
		log.Println(err)
		return nil, err
	}
	out.CreatedAt = timestamppb.New(t)
	return &out, nil
}
