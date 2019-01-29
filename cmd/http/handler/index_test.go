package handler_test

import (
	"net/http"
	"os"
	"testing"

	"git.links123.net/links123.com/pkg/request"
	"git.links123.net/links123.com/stats/cmd/http/router"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var handler *gin.Engine

func TestMain(m *testing.M) {
	handler = router.BuildRouter()

	os.Exit(m.Run())
}
func TestHealthy(t *testing.T) {
	resp := request.AssertHttpRequest(handler, request.HttpRequestConfig{
		Method: http.MethodOptions,
		URL:    "/v1/hello",
	})

	assert.Equal(t, http.StatusOK, resp.Code)
}
