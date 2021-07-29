package global

import (
	"fmt"
	"time"
)

type code struct {
	Code string
	Ch   chan string
	T    *time.Timer
}

var VerifyCode code = code{}

func init() {
	VerifyCode = code{
		Ch: make(chan string),
		T:  time.NewTimer(1 * time.Microsecond),
	}

	go func() {
		for {
			select {
			case VerifyCode.Code = <-VerifyCode.Ch:
			case <-VerifyCode.T.C:
				VerifyCode.Code = ""
				fmt.Println("timed out")
			}
		}
	}()
}
