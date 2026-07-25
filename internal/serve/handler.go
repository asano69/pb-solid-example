package serve

import (
	"net/http"

	"github.com/asano69/myapp/internal/assets"
	"github.com/pocketbase/pocketbase/core"
)

// serveShell serves the SPA's static index.html shell. Solid Router decides
// which screen to render client-side, so every route serves the same shell.
// This is left unauthenticated on purpose: it's an empty HTML/JS bundle with
// no data in it, so an unauthenticated visitor only ever sees the login
// screen the SPA renders client-side.
func serveShell(re *core.RequestEvent) error {
	re.Response.Header().Set("Content-Type", "text/html; charset=utf-8")
	http.ServeFileFS(re.Response, re.Request, assets.FS, "index.html")
	return nil
}

// serveFavicon serves Vite's public/favicon.svg, which is copied to the
// root of the build output, so it's served directly rather than under
// /assets/.
func serveFavicon(re *core.RequestEvent) error {
	re.Response.Header().Set("Content-Type", "image/svg+xml")
	http.ServeFileFS(re.Response, re.Request, assets.FS, "favicon.svg")
	return nil
}
