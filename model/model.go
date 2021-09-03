package model

import (
	"database/sql/driver"
	"fmt"
	"strings"
	"sync"
	"time"
)

type BaseModel struct {
	Id        uint64  `gorm:"primary_key;AUTO_INCREMENT;column:id" json:"-"`
	CreatedAt MyTime  `gorm:"column:createdAt" json:"-"`
	UpdatedAt MyTime  `gorm:"column:updatedAt" json:"-"`
	DeletedAt *MyTime `gorm:"column:deletedAt" sql:"index" json:"-"`
}

type MyTime struct {
	time.Time
}

func (t *MyTime) UnmarshalJSON(data []byte) error {

	if string(data) == "null" {

		return nil
	}
	var err error
	//前端接收的时间字符串
	str := string(data)
	//去除接收的str收尾多余的"
	timeStr := strings.Trim(str, "\"")
	t1, err := time.Parse("2006-01-02 15:04:05", timeStr)
	*t = MyTime{t1}
	return err
}

func (t MyTime) MarshalJSON() ([]byte, error) {
	formatted := fmt.Sprintf("\"%s\"", t.Format("2006-01-02 15:04:05"))
	return []byte(formatted), nil
}

func (t MyTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	if t.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return t.Time, nil
}

func (t *MyTime) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = MyTime{Time: value}
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}

type UserInfo struct {
	Id        uint64 `json:"id"`
	Username  string `json:"username"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

type UserList struct {
	Lock  *sync.Mutex
	IdMap map[uint64]*UserInfo
}

type Token struct {
	Token string `json:"token"`
}
