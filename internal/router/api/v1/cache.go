package v1

import (
	"net/http"
	"webconsole/internal/dao/webcache"

	"github.com/gin-gonic/gin"
)

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
func CacheCheck(c *gin.Context) {
	key := c.GetString("userUnit") + c.Param("key")
	if key == "" {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	webcache.CacheCheck(key)

	//respcode.ResponseSuccess(c, string(b))
}

//func (s *Server) DeleteHandler(c *gin.Context) {
//	key := c.Param("key")
//
//	if key == "" {
//		c.JSON(http.StatusBadRequest, nil)
//		return
//	}
//
//	err := s.Del(key)
//	if err != nil {
//		// 缓存更新失败后 需要加入消息队列重试
//		log.Println(err)
//	}
//}
//
//func (s *Server) StatusHandler(c *gin.Context) {
//	log.Println(s.GetStat())
//	b, err := json.Marshal(s.GetStat())
//	if err != nil {
//		log.Println(err)
//		c.JSON(http.StatusInternalServerError, nil)
//		return
//	}
//
//	respcode.ResponseSuccess(c, string(b))
//}
