package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type API struct {
}

func NewAPI() *API {
	return &API{}
}

func handleError(c *gin.Context, code int, err error) {

	res := ErrorResponse{
		Code:    code,
		Message: err.Error(),
	}
	c.JSON(code, res)
}

func handleRequestError(c *gin.Context, err error) {
	handleError(c, http.StatusBadRequest, err)
}

func handleServerError(c *gin.Context, err error) {
	handleError(c, http.StatusInternalServerError, err)
}
