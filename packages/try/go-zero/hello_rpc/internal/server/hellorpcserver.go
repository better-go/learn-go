// Code generated by goctl. DO NOT EDIT!
// Source: hello_rpc.proto

package server

import (
	"context"

	"zero/hello_rpc/hello_rpc"
	"zero/hello_rpc/internal/logic"
	"zero/hello_rpc/internal/svc"
)

type HelloRpcServer struct {
	svcCtx *svc.ServiceContext
	hello_rpc.UnimplementedHelloRpcServer
}

func NewHelloRpcServer(svcCtx *svc.ServiceContext) *HelloRpcServer {
	return &HelloRpcServer{
		svcCtx: svcCtx,
	}
}

func (s *HelloRpcServer) Ping(ctx context.Context, in *hello_rpc.Request) (*hello_rpc.Response, error) {
	l := logic.NewPingLogic(ctx, s.svcCtx)
	return l.Ping(in)
}