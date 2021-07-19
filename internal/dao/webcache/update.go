package webcache

import (
	"fmt"
	"log"
	"net"
	"strconv"
	"webconsole/global"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func UpdataCache(c *gin.Context, key string, info string) {
	key = c.GetString("userUnit") + key

	klen := strconv.Itoa(len(key))
	vlen := strconv.Itoa(len(info))
	test := "S" + klen + " " + vlen + " " + key + info

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

func DeleteCache(c *gin.Context, key string, info string) {
	key = c.GetString("userUnit") + key
	log.Println(key)

	klen := strconv.Itoa(len(key))
	vlen := strconv.Itoa(len(info))
	test := "S" + klen + " " + vlen + " " + key + info

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

}
