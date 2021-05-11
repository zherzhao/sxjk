package v1

import (
	"encoding/json"
	"log"
	"net/http"
	"webconsole/pkg/respcode"

	"webconsole/pkg/cache/ICache"

	"github.com/gin-gonic/gin"
)

type Server struct {
	ICache.Cache
}

func NewServer(c ICache.Cache) *Server {
	return &Server{c}
}

// CacheCheck 获取缓存数据接口
// @Summary 检查缓存命中接口
// @Description 检查缓存中是否有请求的值 有就返回没有将请求转发
// @Tags 缓存相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param infotype path string true "查询类型 : road(路)  bridge(桥) tunnel(隧道) service(服务区) portal(收费门架) toll(收费站)"
// @Param level path int true "查询等级 : 0(高速) 1(一级) 2(二级) 3(三级) 4(四级) 5(等外)"
// @Security ApiKeyAuth
// @Success 200 {object} respcode.ResponseData{msg=string,data=string}
// @Router /api/v1/cache/hit/{infotype}/{level} [get]
func (s *Server) CacheCheck(c *gin.Context) {
	key := c.Param("key")
	if key == "" {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	b, _ := s.Get(key)
	if len(b) == 0 {
		c.Set("miss", true) // 需要查数据库
		return
	}

	respcode.ResponseSuccess(c, string(b))
	c.Set("miss", false) // 不需要查数据库

}

func (s *Server) StatusHandler(c *gin.Context) {
	log.Println(s.GetStat())
	b, err := json.Marshal(s.GetStat())
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	respcode.ResponseSuccess(c, string(b))
}
