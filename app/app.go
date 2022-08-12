package app

import (
	"embed"
	"io/fs"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/ikmski/go-vue-sample/api"
)

type Config struct {
	Assets embed.FS
}

func NewApp(config *Config) *gin.Engine {

	//var err error

	engine := gin.Default()

	// Render
	engine.Delims("[[", "]]")
	engine.LoadHTMLFiles("frontend/dist/index.html")

	// Static files
	engine.StaticFS("assets/", newStaticFS(config.Assets, "frontend/dist/assets/"))

	// Index
	{
		group := engine.Group("/")
		group.Use(noCache())
		handler := newHandler()
		group.GET("/", handler.handle)
	}

	// API
	{
		group := engine.Group("/api")
		group.Use(noCache())

		api := api.NewAPI()

		group.GET("/message", api.GetMessage)
	}

	return engine
}

type staticFS struct {
	assets embed.FS
	root   string
}

func newStaticFS(assets embed.FS, root string) http.FileSystem {

	fs := &staticFS{
		assets: assets,
		root:   root,
	}

	return http.FS(fs)
}

func (fs *staticFS) Open(name string) (fs.File, error) {

	path := filepath.Join(fs.root, name)
	return fs.assets.Open(path)
}
