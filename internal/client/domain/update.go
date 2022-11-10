package domain

import (
	"configapi/pb"
	"context"
)

func (c *Client) Update(ctx context.Context) error {
	output, err := c.Repo.ReadConfigFromFile()
	if err != nil {
		return err
	}
	updateConfig := string(output)
	response, err := c.ClientProtobuf.Update(ctx,
		&pb.UpdateConfig{Id: c.Repo.GetID(), Config: &updateConfig})
	if err != nil {
		return err
	}
	if err := printConf("Update", response); err != nil {
		return err
	}
	return nil
}
