package v1

import (
	"log"
	"webconsole/internal/dao/database"
	"webconsole/pkg/respcode"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	users, err := database.GetUsersHandler()
	if err != nil {
		log.Println(err)
		respcode.ResponseErrorWithMsg(c, respcode.CodeServerBusy, err)
		return
	}

	respcode.ResponseSuccess(c, users)
}
