package router

import (
	"git.links123.net/links123.com/stats/config"
	"github.com/gin-gonic/gin"
)

var r = gin.Default()

// BuildRouter gin router
func BuildRouter() *gin.Engine {
	// gin config
	r.RedirectTrailingSlash = true
	r.RedirectFixedPath = true

	if config.C.App.Debug {
		// set gin debug mode
		gin.SetMode(gin.DebugMode)
	}

	registerV1Router()

	return r
}
