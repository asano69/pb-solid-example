// Package serve implements the "serve" command, which runs a single HTTP
// server that hosts the SPA shell and static assets.
//
// The package is split across files:
//   - serve.go:   route registration and server startup (this file)
//   - handler.go: HTTP handlers
package serve

import (
	"fmt"
	"io/fs"

	"github.com/asano69/myapp/internal/assets"
	"github.com/asano69/myapp/internal/config"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"

	"github.com/sirupsen/logrus"
)

// Run registers all routes and starts listening.
func Run(app *pocketbase.PocketBase, cfg *config.Config) error {

	// assetsFS exposes just the "assets/" subdirectory that Vite's default
	// (unprefixed) base writes hashed JS/CSS bundles into, so they're served
	// at the conventional /assets/... URL instead of /static/assets/....
	assetsFS, err := fs.Sub(assets.FS, "assets")
	if err != nil {
		return fmt.Errorf("sub assets fs: %w", err)
	}
	addr := fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port)

	app.OnServe().BindFunc(func(e *core.ServeEvent) error {
		e.Router.GET("/assets/{path...}", apis.Static(assetsFS, false))
		e.Router.GET("/", serveShell)
		e.Router.GET("/favicon.svg", serveFavicon)
		return e.Next()
	})

	logrus.WithField("addr", addr).Info("listening")
	return apis.Serve(app, apis.ServeConfig{
		HttpAddr:        addr,
		ShowStartBanner: false,
	})
}
