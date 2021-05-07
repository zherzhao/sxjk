package webcache

import (
	"fmt"
	"net"
	"strconv"
	"webconsole/global"

	"go.uber.org/zap"
)

func UpdataCache(key string, info string) {
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
		zap.L().Error("Write to server failed: ", zap.String("", err.Error()))
		return
	}

	reply := make([]byte, 1024)
	_, err = conn.Read(reply)
	if err != nil {
		zap.L().Error("Write to server failed: ", zap.String("", err.Error()))
		return
	}

	zap.L().Info("reply from server: ", zap.String("", string(reply)))

}
