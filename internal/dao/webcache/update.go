package webcache

import (
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"webconsole/global"
)

func UpdataCache(key string, info string) {
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
	if err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}

	defer conn.Close()

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

}
