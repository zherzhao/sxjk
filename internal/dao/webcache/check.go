package webcache

import (
	"fmt"
	"log"
	"net"
	"strconv"
	"webconsole/global"

	"go.uber.org/zap"
)

func CacheCheck(key string) {

	klen := strconv.Itoa(len(key))
	test := "G" + klen + " " + key

	serverAddr := fmt.Sprintf("127.0.0.1:%s", global.CacheSetting.Port)
	tcpAddr, err := net.ResolveTCPAddr("tcp", serverAddr)
	if err != nil {
		zap.L().Error("ResolveTCPAddr to server failed: ", zap.String("", err.Error()))
		return
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		zap.L().Error("Dial to server failed: ", zap.String("", err.Error()))
		return
	}

	defer conn.Close()

	_, err = conn.Write([]byte(test))
	if err != nil {
		// 缓存更新失败后 需要加入消息队列重试
		zap.L().Error("Write to cache server failed: ", zap.String("", err.Error()))
		return
	}

	reply := make([]byte, 1)
	_, err = conn.Read(reply)
	if err != nil {
		zap.L().Error("Read from cache server failed: ", zap.String("", err.Error()))
		return
	}
	log.Println("test reply:", string(reply))

}

//func DeleteHandler(c *gin.Context) {
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
