package initilization

import (
	"cloudstore/pkg/global"
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() {
	// 初始化配置结构体
	viper.AddConfigPath("config")
	viper.SetConfigName("mysql")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	global.MYSQLCONFIG.Host = viper.GetString("db1.host")
	global.MYSQLCONFIG.Port = viper.GetInt("db1.port")
	global.MYSQLCONFIG.UserName = viper.GetString("db1.username")
	global.MYSQLCONFIG.Password = viper.GetString("db1.password")
	global.MYSQLCONFIG.DBName = viper.GetString("db1.dbname")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		global.MYSQLCONFIG.UserName,
		global.MYSQLCONFIG.Password,
		global.MYSQLCONFIG.Host,
		global.MYSQLCONFIG.Port,
		global.MYSQLCONFIG.DBName,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	global.MYSQLDB = db
}
