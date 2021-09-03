package user

import (
	. "APISERVER/handler"
	"APISERVER/model"
	"APISERVER/pkg/email"
	"APISERVER/pkg/errno"
	"APISERVER/util"
	"github.com/gin-gonic/gin"
	"github.com/zxmrlc/log"
	"github.com/zxmrlc/log/lager"
	"net"
)

// Register @Summary Add new user to the database
// @Description Add a new user
// @Tags user
// @Accept  json
// @Produce  json
// @Param user body user.CreateRequest true "Create a new user"
// @Success 200 {object} user.CreateResponse "{"code":0,"message":"OK","data":{"username":"kong"}}"
// @Router /user [post]
func Register(c *gin.Context) {
	log.Info("User Create function called.", lager.Data{"X-Request-Id": util.GetReqID(c)})
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

	u := model.UserModel{
		Email:      r.Email,
		Password:   r.Password,
		ClientIP:   ip,
		ClientPort: port,
	}

	// Validate the data.
	if err := u.Validate(); err != nil {
		SendResponse(c, errno.ErrValidation, nil)
		return
	}

	// Encrypt the user password
	if err := u.Encrypt(); err != nil {
		SendResponse(c, errno.ErrEncrypt, nil)
		return
	}

	// Insert the user to the database.
	if err := u.Create(); err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	rsp := RegisterResponse{
		Email: r.Email,
	}

	email.SendActivate()

	SendResponse(c, nil, rsp)
}
