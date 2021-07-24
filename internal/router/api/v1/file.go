package v1

import (
	"errors"
	"io/ioutil"
	"os"
	"path"

	"webconsole/pkg/respcode"

	"github.com/gin-gonic/gin"
)

func UploadTable(c *gin.Context) {
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		respcode.ResponseErrorWithMsg(c, respcode.CodeServerBusy, errors.New("文件上传失败"))
		c.Abort()
		return
	}

	content, err := ioutil.ReadAll(file)
	if err != nil {
		respcode.ResponseErrorWithMsg(c, respcode.CodeServerBusy, errors.New("文件读取失败"))
		c.Abort()
		return
	}
	fp := path.Join("/tmp", header.Filename)
	os.WriteFile(fp, content, 0644)
	c.Set("tableName", fp)

	respcode.ResponseSuccess(c, nil)
	c.Next()

}
