package testgrp

import (
	"context"
	"errors"
	"github.com/theo-bot/service4.1-video/foundation/web"
	"math/rand"
	"net/http"
)

// Test is our example route
func Test(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	if n := rand.Intn(100); n%2 == 0 {
		return errors.New("UNTRUSTED ERROR")
	}

	status := struct {
		Status string
	}{
		Status: "OK",
	}

	return web.Respond(ctx, w, status, http.StatusOK)
}
