package user

import (
	. "APISERVER/handler"
	"APISERVER/model"
	"APISERVER/pkg/errno"
	"APISERVER/pkg/token"
	"github.com/gin-gonic/gin"
)

func BindPrinter(c *gin.Context) {
	ctx, err := token.ParseRequest(c)
	if err != nil {
		SendResponse(c, err, nil)
		return
	}

	ps := make([]model.PrinterModel, 2)
	if err := c.Bind(&ps); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}

	u, err := model.GetUser(ctx.Username)
	if err != nil {
		SendResponse(c, errno.ErrUserNotFound, nil)
		return
	}

	for _, printer := range ps {
		p, err := model.GetPrinter(printer.Uuid)
		if err != nil {
			SendResponse(c, err, nil)
			return
		}
		//model.DB.Self.
		if d := model.DB.Self.Model(&u).Association("Printers").Append(p); d.Error != nil {
			SendResponse(c, errno.ErrDatabase, nil)
			return
		}
	}
	SendResponse(c, nil, nil)
}
