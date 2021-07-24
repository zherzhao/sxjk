package v2

import (
	"net/http"
	"net/http/httputil"
	"strconv"

	"github.com/gin-gonic/gin"
)

const Host = "59.49.106.69:8090"

// IServerHandler 获取IServer数据接口
// @Summary 获取IServer数据接口
// @Description 用于转发IServer数据，并做权限检验
// @Tags IServer相关api
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param 查询信息 body IServerReq true "查询组建构建信息"
// @Success 200 {object} model.IServerReq{}
// @Security ApiKeyAuth
// @Success 200 {object} model.IServerResp{}
// @Router /api/v2/iserver/services/{服务}/rest/{服务} [post]
func IServerHandler(ctx *gin.Context) {
	if len(ctx.Request.URL.Path) < 7 {
		return
	}
	ctx.Request.URL.Path = ctx.Request.URL.Path[7:]
	var simpleHostProxy = httputil.ReverseProxy{
		Director: func(req *http.Request) {
			req.URL.Scheme = "http"
			req.URL.Host = Host
			req.Host = Host
			Len, _ := strconv.Atoi(ctx.Request.Header.Get("Content-Length"))
			req.ContentLength = int64(Len)
		},
	}
	simpleHostProxy.ServeHTTP(ctx.Writer, ctx.Request)
}
