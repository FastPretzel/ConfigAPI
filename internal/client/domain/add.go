package domain

import (
	"configapi/pb"
	"context"
	"fmt"
)

func (c *Client) Add(ctx context.Context) error {
	config, err := c.Repo.ReadConfigFromFile()
	if err != nil {
		return err
	}
	idx, err := c.ClientProtobuf.Add(ctx, &pb.Config{Service: *c.Repo.ServiceF,
		Config: string(config)})
	if err != nil {
		return err
	}
	fmt.Printf("Add response: %v\n", idx.GetValue())
	return nil
}
