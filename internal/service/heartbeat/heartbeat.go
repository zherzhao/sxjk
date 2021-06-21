package heartbeat

import (
	"fmt"
	"webconsole/pkg/zinx/ziface"
	"webconsole/pkg/zinx/znet"
)

type PingRouter struct {
	znet.BaseRouter
}

type HelloRouter struct {
	znet.BaseRouter
}

const (
	MsgPing  uint32 = 0
	MsgHello uint32 = 1
)

// Ping Handle
func (p *PingRouter) Handle(request ziface.IRequest) {
	fmt.Println("recv from client:msgID = ", request.GetMsgID(),
		",data = ", string(request.GetMsgData()))

	// 先读取客户端数据 再写回数据到客户端
	err := request.GetConnection().SendMsg(222, []byte("pong...pong..pong.\n"))
	if err != nil {
		fmt.Println("call back error")
	}
}

func (h *HelloRouter) Handle(request ziface.IRequest) {
	fmt.Println("recv from client:msgID = ", request.GetMsgID(),
		",data = ", string(request.GetMsgData()))

	// 先读取客户端数据 再写回数据到客户端
	err := request.GetConnection().SendMsg(666, []byte("hello...hello..hello.\n"))
	if err != nil {
		fmt.Println("call back error")
	}
}
