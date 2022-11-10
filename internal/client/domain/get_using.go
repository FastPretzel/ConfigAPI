package domain

import (
	"configapi/pb"
	"context"
)

func (c *Client) GetUsingConf(ctx context.Context) error {
	config, err := c.ClientProtobuf.GetUsingConf(ctx, &pb.Service{Service: c.Repo.GetService()})
	if err != nil {
		return err
	}
	if err := printConf("GetUsingConf", config); err != nil {
		return err
	}
	return nil
}
