package handlers

import (
	"github.com/dimfeld/httptreemux/v5"
	"github.com/theo-bot/service4.1-video/app/services/sales-api/handlers/v1/testgrp"
	"go.uber.org/zap"
	"net/http"
	"os"
)

// APIMuxConfig contains all the mandatory systems requirements by handlers
type APIMuxConfig struct {
	Shutdown chan os.Signal
	Log      *zap.SugaredLogger
}

// APIMux construcs a http.Handler with all application routers defined
func APIMux(cfg APIMuxConfig) *httptreemux.ContextMux {
	mux := httptreemux.NewContextMux()

	mux.Handle(http.MethodGet, "/test", testgrp.Test)

	return mux
}
