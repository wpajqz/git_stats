package handler

import (
	"git.links123.net/links123.com/pkg/request"
	"git.links123.net/links123.com/stats/service/api"
	"github.com/Unknwon/i18n"
	"github.com/gin-gonic/gin"
)

// Healthy health check for gin
func Healthy(ctx *gin.Context) {
	cp := &api.CommonParams{}
	if request.ParseParamFail(ctx, cp) {
		return
	}

	request.Success(ctx, gin.H{"msg": i18n.Tr(cp.Lang, "hi", "Paul")})
}
