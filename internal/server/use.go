package server

import (
	"configapi/pb"
	"context"
	"log"
)

func (s *Server) Use(ctx context.Context, in *pb.ConfigID) (*pb.ConfigResponse, error) {
	arg := in.GetValue()
	if _, err := s.pool.Exec(ctx, updateInUseQuery, arg); err != nil {
		log.Println(err)
		return nil, err
	}
	return s.getConf(ctx, selectConfigByIdQuery, arg)
}
