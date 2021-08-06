package v2

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"strconv"
	"webconsole/global"
	"webconsole/internal/model"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// IServerHandler 获取IServer数据接口
// @Summary 获取IServer数据接口
// @Description 用于转发IServer数据，并做权限检验
// @Tags IServer相关api
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param 查询信息 body model.IServerReq true "查询组建构建信息"
// @Success 200 {object} model.IServerReq{}
// @Security ApiKeyAuth
// @Success 200 {object} model.IServerResp{}
// @Router /api/v2/iserver/services/{service}/rest/{service} [post]
func IServerPostHandler(ctx *gin.Context) {
	if len(ctx.Request.URL.Path) < 7 {
		return
	}
	ctx.Request.URL.Path = ctx.Request.URL.Path[7:]
	var simpleHostProxy = httputil.ReverseProxy{
		Director: func(req *http.Request) {
			req.URL.Scheme = "http"
			req.URL.Host = global.ServerSetting.Host
			req.Host = global.ServerSetting.Host
			Len, _ := strconv.Atoi(ctx.Request.Header.Get("Content-Length"))
			req.ContentLength = int64(Len)
		},
		// 自定义ModifyResponse
		ModifyResponse: func(resp *http.Response) error {
			if ctx.GetString("userUnit") == "交科" || resp.Request.Method != "POST" {
				return nil
			}

			var oldData, newData []byte

			gr, err := gzip.NewReader(resp.Body)
			if err != nil {
				return err
			}
			defer gr.Close()
			data, err := ioutil.ReadAll(gr)
			if err != nil {
				return err
			}

			if resp.StatusCode < 443 {
				test := new(model.IServerFeatures)
				json.Unmarshal(data, test)
				for i, v := range test.Features {
					for k, name := range v.FieldNames {
						if name == "PTX" || name == "PTY" ||
							name == "经度" || name == "纬度" {
							test.Features[i].FieldNames = test.Features[i].FieldNames[:k-1]
							test.Features[i].FieldValues = test.Features[i].FieldValues[:k-1]
						}
					}
				}

				oldData, _ = json.Marshal(test)
				var b bytes.Buffer
				gz := gzip.NewWriter(&b)
				if _, err := gz.Write(oldData); err != nil {
					zap.L().Error("压缩错误", zap.Error(err))
				}
				if err := gz.Flush(); err != nil {
					zap.L().Error("压缩错误", zap.Error(err))
				}
				if err := gz.Close(); err != nil {
					zap.L().Error("压缩错误", zap.Error(err))
				}

				newData = b.Bytes()
			} else {
				newData, _ = ioutil.ReadAll(resp.Body)
			}
			// 修改返回内容及ContentLength
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(newData))
			resp.ContentLength = int64(len(newData))
			resp.Header.Set("Content-Length", fmt.Sprint(len(newData)))
			return nil
		},
	}
	simpleHostProxy.ServeHTTP(ctx.Writer, ctx.Request)
}
