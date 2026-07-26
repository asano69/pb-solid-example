// Package serve implements the "serve" command, which runs a single HTTP
// server that hosts the SPA shell and static assets.
//
// The package is split across files:
//   - serve.go:   route registration and server startup (this file)
//   - handler.go: HTTP handlers
package serve

import (
	"fmt"
	"log/slog"

	"github.com/asano69/myapp/internal/assets"
	"github.com/asano69/myapp/internal/config"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
)

// Run registers all routes and starts listening.
func Run(app *pocketbase.PocketBase, cfg *config.Config) error {
	addr := fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port)
	app.OnServe().BindFunc(func(e *core.ServeEvent) error {
		// Serves the whole Vite build output (index.html, hashed JS/CSS
		// under assets/, and public/ files like favicon.svg copied to the
		// root) from a single route. indexFallback=true makes any unmatched
		// path fall back to index.html, so client-side routing still works
		// on a hard refresh.
		e.Router.GET("/{path...}", apis.Static(assets.FS, true))
		return e.Next()
	})

	slog.Info("listening", "addr", addr)
	return apis.Serve(app, apis.ServeConfig{
		HttpAddr:        addr,
		ShowStartBanner: false,
	})
}
