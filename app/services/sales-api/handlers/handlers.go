package handlers

import (
	"github.com/theo-bot/service4.1-video/app/services/sales-api/handlers/v1/testgrp"
	"github.com/theo-bot/service4.1-video/business/web/v1/mid"
	"github.com/theo-bot/service4.1-video/foundation/web"
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
func APIMux(cfg APIMuxConfig) *web.App {
	app := web.NewApp(cfg.Shutdown, mid.Logger(cfg.Log), mid.Errors(cfg.Log), mid.Panics())

	app.Handle(http.MethodGet, "/test", testgrp.Test)

	return app
}
