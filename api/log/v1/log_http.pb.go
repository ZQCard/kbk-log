// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.5.2
// - protoc             v4.23.0--rc1
// source: log/v1/log.proto

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationLogServiceCreateLog = "/log.v1.LogService/CreateLog"
const OperationLogServiceGetLogList = "/log.v1.LogService/GetLogList"

type LogServiceHTTPServer interface {
	CreateLog(context.Context, *CreateLogReq) (*emptypb.Empty, error)
	GetLogList(context.Context, *GetLogListReq) (*GetLogListPageRes, error)
}

func RegisterLogServiceHTTPServer(s *http.Server, srv LogServiceHTTPServer) {
	r := s.Route("/")
	r.GET("/log", _LogService_GetLogList0_HTTP_Handler(srv))
	r.POST("/log", _LogService_CreateLog0_HTTP_Handler(srv))
}

func _LogService_GetLogList0_HTTP_Handler(srv LogServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetLogListReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationLogServiceGetLogList)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetLogList(ctx, req.(*GetLogListReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetLogListPageRes)
		return ctx.Result(200, reply)
	}
}

func _LogService_CreateLog0_HTTP_Handler(srv LogServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateLogReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationLogServiceCreateLog)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateLog(ctx, req.(*CreateLogReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*emptypb.Empty)
		return ctx.Result(200, reply)
	}
}

type LogServiceHTTPClient interface {
	CreateLog(ctx context.Context, req *CreateLogReq, opts ...http.CallOption) (rsp *emptypb.Empty, err error)
	GetLogList(ctx context.Context, req *GetLogListReq, opts ...http.CallOption) (rsp *GetLogListPageRes, err error)
}

type LogServiceHTTPClientImpl struct {
	cc *http.Client
}

func NewLogServiceHTTPClient(client *http.Client) LogServiceHTTPClient {
	return &LogServiceHTTPClientImpl{client}
}

func (c *LogServiceHTTPClientImpl) CreateLog(ctx context.Context, in *CreateLogReq, opts ...http.CallOption) (*emptypb.Empty, error) {
	var out emptypb.Empty
	pattern := "/log"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationLogServiceCreateLog))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *LogServiceHTTPClientImpl) GetLogList(ctx context.Context, in *GetLogListReq, opts ...http.CallOption) (*GetLogListPageRes, error) {
	var out GetLogListPageRes
	pattern := "/log"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationLogServiceGetLogList))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
