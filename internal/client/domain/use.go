package domain

import (
	"configapi/pb"
	"context"
)

func (c *Client) Use(ctx context.Context) error {
	config, err := c.ClientProtobuf.Use(ctx, &pb.ConfigID{Value: c.Repo.GetID()})
	if err != nil {
		return err
	}
	if err := printConf("Use", config); err != nil {
		return err
	}
	return nil
}
