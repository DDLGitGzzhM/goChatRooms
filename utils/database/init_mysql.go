package database

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var DB *gorm.DB

// InitMySQL
// 初始化MySQL的连接
func InitMySQL() {

	// 自定义日志模板 打印sql语句

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			//慢sql阈值
			SlowThreshold: time.Second,
			//级别
			LogLevel: logger.Info,
			//彩色
			Colorful: true,
		},
	)

	// 启动mysql
	var err error
	DB, err = gorm.Open(mysql.Open(viper.GetString("mysql.dns")),
		&gorm.Config{Logger: newLogger})

	if err != nil {
		log.Fatalln("Mysql init :", err)
		return
	}

	fmt.Println("Mysql inited ...")
}
