package v1

import (
	"net/http"
	"net/http/httputil"

	"github.com/gin-gonic/gin"
)

const Host = "59.49.106.69:8090"

var simpleHostProxy = httputil.ReverseProxy{
	Director: func(req *http.Request) {
		req.URL.Scheme = "http"
		req.URL.Host = Host
		req.URL.Path = "/iserver/services/map-2020/rest/maps/2020"
		req.Host = Host
	},
}

func MapHandler(ctx *gin.Context) {

	ctx.Request.URL.Path = ctx.Request.URL.Path[7:]
	simpleHostProxy.ServeHTTP(ctx.Writer, ctx.Request)
}
