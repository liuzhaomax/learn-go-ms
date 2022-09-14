package internal

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"learn-go-ms/account_srv/model"
	"log"
	"os"
	"time"
)

var DB *gorm.DB
var err error

func InitDB() {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: true,
			Colorful:                  true,
		},
	)
	conn := fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", "root", "123456", "localhost", 3308, "happy_account_mic_training")
	DB, err = gorm.Open(mysql.Open(conn), &gorm.Config{
		Logger: newLogger,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		panic("数据库连接失败：" + err.Error())
	}
	err = DB.AutoMigrate(&model.Account{})
	if err != nil {
		fmt.Println(err)
	}
}

type DBConfig struct {
}
