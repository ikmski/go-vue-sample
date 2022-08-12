package app

import (
	"github.com/gin-gonic/gin"
)

func noCache() gin.HandlerFunc {

	return func(c *gin.Context) {
		c.Header("Cache-Control", "no-cache")
	}
}
