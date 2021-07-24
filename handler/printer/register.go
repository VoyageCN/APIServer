package printer

import (
	. "APISERVER/handler"
	"APISERVER/model"
	"APISERVER/pkg/errno"
	"github.com/gin-gonic/gin"
	"net"
)

func Register(c *gin.Context) {
	var r RegisterRequest

	if err := c.Bind(&r); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}

	ip, port, err := net.SplitHostPort(c.Request.RemoteAddr)
	if err != nil {
		SendResponse(c, err, nil)
		return
	}

	p := &model.PrinterModel{
		Uuid: r.Uuid,
		Host: ip,
		Port: port,
	}

	if err := p.Create(); err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	rsp := RegisterResponse{
		Uuid: p.Uuid,
	}

	SendResponse(c, nil, rsp)
}
