package main

import (
	"embed"

	"github.com/ikmski/go-vue-sample/app"
)

//go:embed frontend/dist
var assets embed.FS

func main() {

	config := newConfig()

	appConfig := &app.Config{
		Assets: assets,
	}

	app := app.NewApp(appConfig)

	if config.server.TLS {
		app.RunTLS(
			config.server.getAddr(),
			config.server.CertFile,
			config.server.KeyFile)
	} else {
		app.Run(config.server.getAddr())
	}
}
