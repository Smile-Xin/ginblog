package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

func Cors() gin.HandlerFunc {

	return cors.New(cors.Config{
		AllowAllOrigins: true,
		//AllowOrigins:           nil,
		//AllowOriginFunc:        nil,
		AllowMethods: []string{"*"},
		AllowHeaders: []string{"Origin", "content-type", "Authorization"},
		//AllowCredentials:       false,
		ExposeHeaders: []string{"Content-Length", "Authorization"},
		MaxAge:        12 * time.Hour,
		//AllowWildcard:          false,
		//AllowBrowserExtensions: false,
		//AllowWebSockets: false,
		//AllowFiles: false,
	})
}
