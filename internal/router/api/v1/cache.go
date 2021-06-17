package v1

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"webconsole/pkg/cache/ICache"
	"webconsole/pkg/respcode"

	"github.com/gin-gonic/gin"
)

type Server struct {
	ICache.Cache
}

func NewServer(c ICache.Cache) *Server {
	return &Server{c}
}

// CacheCheck 检查缓存命中接口
// @Summary 检查缓存命中接口
// @Description 检查缓存中是否有请求的值 有就返回没有将请求转发
// @Tags 缓存相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string false "Bearer 用户令牌"
// @Param infotype query string true "查询类型"
// @Param level query string true "查询等级"
// @Security ApiKeyAuth
// @Success 200 {string} string "成功"
// @Router /api/v1/cache/hit/{infotype}/{year}/{level} [get]
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

func (s *Server) UpdateHandler(c *gin.Context) {
	key := c.Param("key")

	if key == "" {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	b, _ := ioutil.ReadAll(c.Request.Body)
	if len(b) != 0 {
		e := s.Set(key, b)
		if e != nil {
			log.Println(e)
			c.JSON(http.StatusInternalServerError, nil)
		}
	}
}

func (s *Server) DeleteHandler(c *gin.Context) {
	key := c.Param("key")

	if key == "" {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	e := s.Del(key)
	if e != nil {
		log.Println(e)
		c.JSON(http.StatusInternalServerError, nil)
	}
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
