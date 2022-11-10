package domain

import (
	"configapi/pb"
	"context"
	"io"
	"log"
)

func (c *Client) GetAllServiceConf(ctx context.Context) error {
	stream, err := c.ClientProtobuf.GetAllServiceConf(context.Background(),
		&pb.Service{Service: c.Repo.GetService()})
	if err != nil {
		return err
	}
	for {
		config, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("%v.ListFeatures(_) = _, %v", c, err)
			return err
		}
		err = printConf("GetAllServiceConf", config)
	}
	return nil
}
