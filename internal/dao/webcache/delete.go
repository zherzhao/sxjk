package webcache

import (
	"log"

	"github.com/impact-eintr/ecache/client"
)

func CacheDelete(key string) {

	cmd := &client.Cmd{
		Name: "del",
		Key:  key,
	}
	cli, err := client.New("tcp", "127.0.0.1:6430")
	if err != nil {
		log.Println(err)
		return
	}
	cli.Run(cmd)

}
