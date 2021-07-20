package v2

import (
	"net/http"
	"net/http/httputil"

	"github.com/gin-gonic/gin"
)

const Host = "59.49.106.69:8090"

type IServerResp struct {
}

var simpleHostProxy = httputil.ReverseProxy{
	Director: func(req *http.Request) {
		req.URL.Scheme = "http"
		req.URL.Host = Host
		req.Host = Host
	},
}

// IServerHandler 获取IServer数据接口
// @Summary 获取IServer数据接口
// @Description 用于转发IServer数据，并做权限检验
// @Tags IServer相关api
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Security ApiKeyAuth
// @Success 200 {object} IServerResp{}
// @Router /api/v2/iserver/services/data-vector/rest/data [post]
func IServerHandler(ctx *gin.Context) {
	ctx.Request.URL.Path = ctx.Request.URL.Path[7:]
	simpleHostProxy.ServeHTTP(ctx.Writer, ctx.Request)
}
