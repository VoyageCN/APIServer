package printer

import (
	. "APISERVER/handler"
	"APISERVER/model"
	"APISERVER/pkg/errno"
	"github.com/gin-gonic/gin"
)

func Connect(c *gin.Context) {
	p := &model.PrinterModel{}
	if err := c.Bind(&p); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}

}
