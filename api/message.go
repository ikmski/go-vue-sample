package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Message struct {
	Message string `json:"message"`
}

func (api *API) GetMessage(c *gin.Context) {

	message := &Message{
		Message: "HELLO",
	}

	c.JSON(http.StatusOK, message)
}
