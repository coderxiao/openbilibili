// Code generated by protoc-gen-liverpc v0.1, DO NOT EDIT.
// source: v1/SilentMng.proto

package v1

import context "context"

import proto "github.com/golang/protobuf/proto"
import "go-common/library/net/rpc/liverpc"

var _ proto.Message // generate to suppress unused imports

// ===================
// SilentMng Interface
// ===================

type SilentMng interface {
	// * 查询是否是黑名单
	//
	IsBlockUser(context.Context, *SilentMngIsBlockUserReq) (*SilentMngIsBlockUserResp, error)
}

// =========================
// SilentMng Live Rpc Client
// =========================

type silentMngRpcClient struct {
	client *liverpc.Client
}

// NewSilentMngRpcClient creates a Rpc client that implements the SilentMng interface.
// It communicates using Rpc and can be configured with a custom HTTPClient.
func NewSilentMngRpcClient(client *liverpc.Client) SilentMng {
	return &silentMngRpcClient{
		client: client,
	}
}

func (c *silentMngRpcClient) IsBlockUser(ctx context.Context, in *SilentMngIsBlockUserReq) (*SilentMngIsBlockUserResp, error) {
	out := new(SilentMngIsBlockUserResp)
	err := doRpcRequest(ctx, c.client, 1, "SilentMng.is_block_user", in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}
