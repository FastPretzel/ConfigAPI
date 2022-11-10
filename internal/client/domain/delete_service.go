package domain

import (
	"configapi/pb"
	"context"
	"fmt"
)

func (c *Client) DeleteService(ctx context.Context) error {
	response, err := c.ClientProtobuf.DeleteService(ctx,
		&pb.Service{Service: c.Repo.GetService()})
	if err != nil {
		return err
	}
	_, err = fmt.Printf("DeleteService response:\n\tdelete: %v\n", response.GetSuccess())
	if err != nil {
		return err
	}
	return nil
}
