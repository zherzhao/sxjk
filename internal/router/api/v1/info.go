package v1

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"
	"webconsole/global"

	"github.com/gin-gonic/gin"
)

type Info struct {
}

func NewInfo() Info {
	return Info{}
}

func (this *Info) GetUpdateInfo(c *gin.Context) {
	info := c.GetString("info")

	c.JSON(http.StatusOK, info) // 向浏览器返回数据

	key := "/" + c.Param("infotype") + "/" + c.Param("count")

	// 如果是缓存在磁盘中
	if global.CacheSetting.CacheType == "disk" {
		go func(string, string) {
			klen := strconv.Itoa(len(key))
			vlen := strconv.Itoa(len(info))
			test := "S" + klen + " " + vlen + " " + key + info

			serverAddr := fmt.Sprintf("127.0.0.1:%s", global.CacheSetting.Port)
			tcpAddr, err := net.ResolveTCPAddr("tcp", serverAddr)
			if err != nil {
				log.Println(err.Error())
				os.Exit(1)

			}

			conn, err := net.DialTCP("tcp", nil, tcpAddr)
			defer conn.Close()

			if err != nil {
				log.Println(err.Error())
				os.Exit(1)

			}
			_, err = conn.Write([]byte(test))
			if err != nil {
				log.Println("Write to server failed:", err.Error())
				os.Exit(1)

			}

			reply := make([]byte, 1024)
			_, err = conn.Read(reply)
			if err != nil {
				log.Println("Write to server failed:", err.Error())
				os.Exit(1)

			}

			fmt.Printf("reply from server:\n%v\n", string(reply))

		}(key, info)

	} else {
		c.Set("type", "mem")
		c.Request.URL.Path = "/api/v1/cache/update" + key //将请求的URL修改
		c.Request.Method = http.MethodPut
		c.Request.Body = ioutil.NopCloser(bytes.NewReader([]byte(info)))

	}

}
