package domain

import (
	repo "configapi/internal/client/repository"
	"configapi/pb"
)

type Client struct {
	ClientProtobuf pb.ConfigServiceClient
	Repo           *repo.Repository
}
