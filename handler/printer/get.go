package printer

import (
	. "APISERVER/handler"
	"APISERVER/model"
	"APISERVER/pkg/errno"
	"github.com/gin-gonic/gin"
)

func Get(c *gin.Context) {
	uuid := c.Query("uuid")
	if uuid == "" {
		SendResponse(c, errno.ErrParameterInvalid, nil)
		return
	}

	printer, err := model.GetPrinter(uuid)
	if err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	SendResponse(c, nil, printer)
}
