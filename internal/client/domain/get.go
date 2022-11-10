package domain

import (
	"configapi/pb"
	"context"
	"fmt"
	"strings"
)

func (c *Client) Get(ctx context.Context) error {
	config, err := c.ClientProtobuf.Get(ctx, &pb.ConfigID{Value: c.Repo.GetID()})
	if err != nil {
		return err
	}
	if err := printConf("Get", config); err != nil {
		return err
	}
	return nil
}

func printConf(method string, config *pb.ConfigResponse) error {
	_, err := fmt.Printf("%v response:\n\tID: %v,\n\tservice: %v,\n\tconfig: %v,"+
		"\n\tversion: %v,\n\tis in use: %v,\n\tcreated at: %v\n",
		method, config.GetId(), config.GetConfig().Service, strings.TrimSuffix(config.GetConfig().Config, "\n"),
		config.GetVersion(), config.GetInUse(), config.GetCreatedAt().AsTime())
	if err != nil {
		return err
	}
	return nil
}
