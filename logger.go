package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {

		logrus.Infof(
			"%s %s",
			c.Request.Method,
			c.Request.URL.Path,
		)

		c.Next()
	}
}
