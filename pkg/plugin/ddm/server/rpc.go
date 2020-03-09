package server

import (
	"context"

	pb "github.com/kubemove/kubemove/pkg/plugin/proto"
)

const (
	OK  = 0
	ERR = 1
)

func (p *plugin) init(params map[string]string) error {
	req := &pb.InitRequest{
		Params: params,
	}

	_, err := p.Init(context.Background(), req)
	if err != nil {
		return err
	}
	return nil
}

func (p *plugin) syncData(params map[string]string) (string, error) {
	req := &pb.SyncRequest{
		Params: params,
	}

	res, err := p.SyncData(context.Background(), req)
	if err != nil {
		return "", err
	}

	return res.SyncID, nil
}

func (p *plugin) syncStatus(params map[string]string) (int32, error) {
	req := &pb.SyncStatusRequest{
		Params: params,
	}

	res, err := p.SyncStatus(context.Background(), req)
	if err != nil {
		return ERR, err
	}

	return res.Status, nil
}
