package domain

import (
	"configapi/pb"
	"context"
	"fmt"
)

func (c *Client) DeleteConf(ctx context.Context) error {
	response, err := c.ClientProtobuf.DeleteConf(ctx, &pb.ConfigID{Value: c.Repo.GetID()})
	if err != nil {
		return err
	}
	_, err = fmt.Printf("DeleteConf response:\n\tdelete: %v\n", response.GetSuccess())
	if err != nil {
		return err
	}
	return nil
}
