// Package serve implements the "serve" command, which runs a single HTTP server
// that hosts the index page and all drill sessions defined in the config file.
package serve

import (
	"fmt"

	"io/fs"
	"net/http"
	"net/url"
	"regexp"
	"sort"
	"strings"

	"github.com/asano69/myapp/internal/config"
	"github.com/asano69/myapp/internal/db"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"

	"github.com/sirupsen/logrus"
)

// Run opens the database and collection once, registers all drill routes, then
// starts listening. The database and collection are shared across all sessions.
func Run(app *pocketbase.PocketBase, cfg *config.Config) error {
	database, err := db.New(app)
	if err != nil {
		return err
	}

	col, err := collection.Load(cfg.Data.Root, database)
	if err != nil {
		return err
	}

	addr := fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port)
	app.OnServe().BindFunc(func(e *core.ServeEvent) error {

		e.Router.GET("/favicon.svg", func(re *core.RequestEvent) error {
			re.Response.Header().Set("Content-Type", "image/svg+xml")
			http.ServeFileFS(re.Response, re.Request, assets.FS, "favicon.svg")
			return nil
		})

		return e.Next()
	})

	logrus.WithField("addr", addr).Info("listening")
	return apis.Serve(app, apis.ServeConfig{
		HttpAddr:        addr,
		ShowStartBanner: false,
	})

}
