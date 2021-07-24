package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
	"github.com/zxmrlc/log"
)

type Database struct {
	Self   *gorm.DB
	Docker *gorm.DB
}

var DB *Database

func openDB(username, password, addr, name string) *gorm.DB {
	config := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=%t&loc=%s",
		username,
		password,
		addr,
		name,
		true,
		//"Asia/Shanghai",
		"Local")

	db, err := gorm.Open("mysql", config)
	if err != nil {
		log.Errorf(err, "Database connection failed, Database name: %s", name)
	}

	setup(db)

	return db
}

func setup(db *gorm.DB) {
	db.LogMode(viper.GetBool("gormlog"))
	db.DB().SetMaxIdleConns(0)
}

func InitSelfDB() *gorm.DB {
	db := openDB(viper.GetString("db.username"),
		viper.GetString("db.password"),
		viper.GetString("db.addr"),
		viper.GetString("db.name"))

	db.AutoMigrate(&UserModel{}, &PrinterModel{})

	return db
}

func GetSelfDB() *gorm.DB {
	return InitSelfDB()
}

func InitDockerDB() *gorm.DB {
	return openDB(viper.GetString("docker_db.username"),
		viper.GetString("docker_db.password"),
		viper.GetString("docker_db.addr"),
		viper.GetString("docker_db.name"))
}

func GetDockerDB() *gorm.DB {
	return InitDockerDB()
}

func (db *Database) Init() {
	DB = &Database{
		Self: GetSelfDB(),
		//Docker: GetDockerDB(),
	}
}

func (db *Database) Close() {
	DB.Self.Close()
	//DB.Docker.Close()
}
