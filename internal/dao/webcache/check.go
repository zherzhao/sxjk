package webcache

import (
	"log"

	"github.com/impact-eintr/ecache/client"
)

func CacheCheck(key string) []byte {

	cmd := &client.Cmd{
		Name: "get",
		Key:  key,
	}
	cli, err := client.New("tcp", "127.0.0.1:6430")
	if err != nil {
		log.Println(err)
		return nil
	}
	cli.Run(cmd)

	return cmd.Value

}
