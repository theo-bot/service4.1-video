package handlers

import (
	"encoding/json"
	"github.com/dimfeld/httptreemux/v5"
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
func APIMux(cfg APIMuxConfig) http.Handler {
	mux := httptreemux.NewContextMux()

	h := func(w http.ResponseWriter, r *http.Request) {
		status := struct {
			Status string
		}{
			Status: "OK",
		}

		json.NewEncoder(w).Encode(status)
	}

	mux.Handle(http.MethodGet, "/test", h)

	return mux
}
