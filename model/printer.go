package model

import "APISERVER/pkg/errno"

type PrinterModel struct {
	BaseModel
	Uuid string `json:"uuid" gorm:"column:uuid;not null;unique" binding:"required"`
	Host string `json:"host" gorm:"column:host;not null"`
	Port string `json:"port" gorm:"column:port;not null"`
}

func (p *PrinterModel) TableName() string {
	return "tb_printers"
}

func (p *PrinterModel) Create() error {
	return DB.Self.Create(&p).Error
}

func (p *PrinterModel) Update() error {
	return DB.Self.Save(&p).Error
}

func GetPrinter(uuid string) (*PrinterModel, error) {
	p := &PrinterModel{}
	d := DB.Self.Where("uuid = ?", uuid).First(&p)
	return p, d.Error
}

func ConnectPrinter(uuid string) error {
	// 客户端发出连接打印机请求
	// 服务器向打印机要求打印机向客户端发送主动帧打洞

	return errno.ErrBind
}
