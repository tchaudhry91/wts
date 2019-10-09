package svc

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"github.com/pkg/errors"
)

type putRequest ScriptRecord

type putResponse struct {
	Err string `json:"err,omitempty"`
}

func makePutEndpoint(svc WTSService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(putRequest)
		record := ScriptRecord(req)
		err = svc.Put(ctx, record.User, &record)
		if err != nil {
			return putResponse{
				Err: errors.WithMessage(err, "Failed to store record").Error(),
			}, errors.WithMessage(err, "Failed to store the record")
		}
		return putResponse{
			Err: "",
		}, nil
	}
}

type getRequest struct {
	Name string `json:"name,omitempty"`
	User string `json:"user,omitempty"`
}

type getResponse struct {
	Record *ScriptRecord `json:"record,omitempty"`
	Err    string        `json:"err,omitempty"`
}

func makeGetEndpoint(svc WTSService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(getRequest)
		record, err := svc.Get(ctx, req.User, req.Name)
		if err != nil {
			return getResponse{
				Err: err.Error(),
			}, errors.WithMessage(err, "Failed to retrieve the record")
		}
		return getResponse{
			Record: record,
		}, nil
	}
}
