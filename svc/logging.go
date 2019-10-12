package svc

import (
	"context"
	"github.com/go-kit/kit/log"
	"time"
)

// LoggingMiddleware takes a logger as a dependency
// and returns a ServiceMiddleware.
func LoggingMiddleware(next WTSService, logger log.Logger) WTSService {
	return loggingMiddleware{
		logger: logger,
		next:   next,
	}
}

type loggingMiddleware struct {
	logger log.Logger
	next   WTSService
}

func (mw loggingMiddleware) Put(ctx context.Context, user string, rec *ScriptRecord) (err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"method", "put",
			"user", user,
			"public", rec.Public,
			"script", rec.Script.Name,
			"took", time.Since(begin),
		)
	}(time.Now())
	return mw.next.Put(ctx, user, rec)
}

func (mw loggingMiddleware) Get(ctx context.Context, user, name string) (record *ScriptRecord, err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"method", "get",
			"user", user,
			"name", name,
			"took", time.Since(begin),
		)
	}(time.Now())
	return mw.next.Get(ctx, user, name)
}
