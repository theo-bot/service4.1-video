package mid

import (
	"context"
	"github.com/theo-bot/service4.1-video/business/web/metrics"
	"github.com/theo-bot/service4.1-video/foundation/web"
	"net/http"
)

func Metrics() web.Middleware {
	m := func(handler web.Handler) web.Handler {
		h := func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
			ctx = metrics.Set(ctx)

			err := handler(ctx, w, r)

			metrics.AddRequests(ctx)
			metrics.AddGoroutines(ctx)

			if err != nil {
				metrics.AddErrors(ctx)
			}

			return err
		}

		return h
	}

	return m
}
