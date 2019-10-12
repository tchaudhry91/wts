package svc

import (
	"context"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/transport"
	kitgrpc "github.com/go-kit/kit/transport/grpc"
	"github.com/tchaudhry91/wts-service/svc/pb"
	"github.com/tchaudhry91/wts-service/wts"
)

type grpcServer struct {
	get kitgrpc.Handler
	put kitgrpc.Handler
}

// NewGRPCServer makes a set of endpoints available as a gRPC AddServer.
func NewGRPCServer(endpoints Endpoints, logger log.Logger) pb.WTSServer {
	options := []kitgrpc.ServerOption{
		kitgrpc.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
	}

	return &grpcServer{
		get: kitgrpc.NewServer(
			endpoints.Get,
			decodeGRPCGetRequest,
			encodeGRPCGetResponse,
			options...,
		),
		put: kitgrpc.NewServer(
			endpoints.Put,
			decodeGRPCPutRequest,
			encodeGRPCPutResponse,
			options...,
		),
	}
}

func (s *grpcServer) Get(ctx context.Context, req *pb.GetRequest) (*pb.GetResponse, error) {
	_, rep, err := s.get.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.GetResponse), nil
}

func (s *grpcServer) Put(ctx context.Context, req *pb.ScriptRecord) (*pb.PutResponse, error) {
	_, rep, err := s.put.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.PutResponse), nil
}

func decodeGRPCPutRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.ScriptRecord)
	return putRequest{User: req.User, Public: req.Public, Script: scriptFromPb(req.Script)}, nil
}

func decodeGRPCGetRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.GetRequest)
	return getRequest{Name: req.Name, User: req.User}, nil
}

func encodeGRPCPutResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(putResponse)
	return &pb.PutResponse{Err: resp.Err}, nil
}

func encodeGRPCGetResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(getResponse)
	return &pb.GetResponse{Record: pbFromScriptRecord(resp.Record), Err: resp.Err}, nil
}

func scriptFromPb(script *pb.Script) *wts.Script {
	return &wts.Script{
		Name:         script.Name,
		Interpretter: script.Interpretter,
		Script:       script.Script,
		Checksum:     script.Checksum,
	}
}

func pbFromScript(script *wts.Script) *pb.Script {
	return &pb.Script{
		Name:         script.Name,
		Interpretter: script.Interpretter,
		Script:       script.Script,
		Checksum:     script.Checksum,
	}
}

func pbFromScriptRecord(record *ScriptRecord) *pb.ScriptRecord {
	return &pb.ScriptRecord{
		User:   record.User,
		Public: record.Public,
		Script: pbFromScript(record.Script),
	}
}
