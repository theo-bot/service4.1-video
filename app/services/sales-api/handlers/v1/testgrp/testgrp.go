package testgrp

import (
	"context"
	"errors"
	v1 "github.com/theo-bot/service4.1-video/business/web/v1"
	"github.com/theo-bot/service4.1-video/foundation/web"
	"math/rand"
	"net/http"
)

// Test is our example route
func Test(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	if n := rand.Intn(100); n%2 == 0 {
		return v1.NewRequestError(errors.New("TRUSTED ERROR"), http.StatusBadRequest)
	}

	status := struct {
		Status string
	}{
		Status: "OK",
	}

	return web.Respond(ctx, w, status, http.StatusOK)
}
