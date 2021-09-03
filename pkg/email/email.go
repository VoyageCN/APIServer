package email

import (
	"github.com/jordan-wright/email"
	"log"
	"net/smtp"
)

func SendActivate() {
	e := email.NewEmail()
	e.From = "dj <ieshiaonan@163.com>"
	e.To = []string{"745831307@qq.com"}
	e.Subject = "这是主题"
	e.Text = []byte("www.topgoer.com是个不错的go语言中文文档")
	// 设置服务器相关配置
	err := e.Send("smtp.163.com:25", smtp.PlainAuth("", "ieshiaonan@163.com", "YXVUDJYWOXSQISKV", "smtp.163.com"))
	if err != nil {
		log.Fatalln(err.Error())
	}
}
