package v1

import (
	"log"
	"net/http"
	"net/http/httputil"

	"github.com/gin-gonic/gin"
)

const Host = "59.49.106.69:8090"

var simpleHostProxy = httputil.ReverseProxy{
	Director: func(req *http.Request) {
		req.URL.Scheme = "http"
		req.URL.Host = Host
		req.Host = Host
	},
}

func DataHandler(ctx *gin.Context) {
	ctx.Request.URL.Path = ctx.Request.URL.Path[7:]
	test := ctx.Request.Header.Get("Authorization")
	log.Println(test)
	simpleHostProxy.ServeHTTP(ctx.Writer, ctx.Request)
}
