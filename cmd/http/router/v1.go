package router

import (
	"git.links123.net/links123.com/stats/cmd/http/handler"
)

func registerV1Router() {
	v1NoAuth := r.Group("/v1")
	{
		v1NoAuth.OPTIONS("/hello", handler.Healthy)
	}
}
