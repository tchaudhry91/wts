package svc

import (
	"context"
)

type WTSService interface {
	Put(ctx context.Context, user string, rec *ScriptRecord) error
	Get(ctx context.Context, user string, name string) (rec *ScriptRecord, err error)
}

type wtsService struct {
	S Store
}

// Put saves the desired script record
func (svc *wtsService) Put(ctx context.Context, user string, rec *ScriptRecord) error {
	return svc.S.Put(ctx, user, rec)
}

// Get retrieves the desired script record
func (svc *wtsService) Get(ctx context.Context, user, name string) (rec *ScriptRecord, err error) {
	return svc.S.Get(ctx, user, name)
}
