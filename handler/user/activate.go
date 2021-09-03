package user

import (
	. "APISERVER/handler"
	"github.com/gin-gonic/gin"
)

func Activate(c *gin.Context) {
	token := c.Query("token")
	// 查表判断是谁的token
	if token == "" {
		return
	}

	// 激活成功
	SendResponse(c, nil, nil)
}
