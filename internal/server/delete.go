package server

import (
	"configapi/pb"
	"context"
	"log"
)

const deleteConfQuery = `DELETE FROM config WHERE config_id = $1`

func (s *Server) DeleteConf(ctx context.Context, in *pb.ConfigID) (*pb.DeleteResponse, error) {
	return s.deleteQuery(ctx, deleteConfQuery, in.GetValue())
}

const deleteServiceQuery = `DELETE FROM service WHERE name = $1`

func (s *Server) DeleteService(ctx context.Context, in *pb.Service) (*pb.DeleteResponse, error) {
	return s.deleteQuery(ctx, deleteServiceQuery, in.GetService())
}

func (s *Server) deleteQuery(ctx context.Context, query string, arg any) (*pb.DeleteResponse, error) {
	res, err := s.pool.Exec(ctx, query, arg)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	cnt := res.RowsAffected()
	if cnt == 0 {
		return &pb.DeleteResponse{Success: false}, nil
	}
	return &pb.DeleteResponse{Success: true}, nil
}
