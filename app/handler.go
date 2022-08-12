package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type handler struct {
}

func newHandler() *handler {
	return &handler{}
}

func (h handler) handle(c *gin.Context) {

	data := getData(c)
	c.HTML(
		http.StatusOK,
		"index.html",
		data)
}

func getData(c *gin.Context) map[interface{}]interface{} {

	d := make(map[interface{}]interface{})

	return d
}
