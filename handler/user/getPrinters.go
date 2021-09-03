package user

import (
	. "APISERVER/handler"
	"APISERVER/model"
	"APISERVER/pkg/errno"
	"APISERVER/pkg/token"
	"github.com/gin-gonic/gin"
)

// GET
func GetPrinters(c *gin.Context) {
	ctx, _ := token.ParseRequest(c)
	if ctx.Username != c.Query("username") {
		SendResponse(c, errno.ErrTokenInvalid, nil)
		return
	}

	u, err := model.GetUser(ctx.Username)
	if err != nil {
		SendResponse(c, errno.ErrUserNotFound, nil)
		return
	}
	SendResponse(c, nil, u.Printers)
}
