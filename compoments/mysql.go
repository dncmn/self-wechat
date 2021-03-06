package compoments

import (
	// 导入mysql驱动
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"self-wechat/config"
	"self-wechat/utils/logging"
)

var (
	dns = fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.Config.Mysql.Username,
		config.Config.Mysql.Password,
		config.Config.Mysql.Host,
		config.Config.Mysql.Dbname,
	)
	globDb *gorm.DB
	logs   = logging.GetLogger()
)

// 连接数据库
//func init() {
//
//	var err error
//	if globDb, err = gorm.Open("mysql", dns); err != nil {
//		log.Fatal(err)
//	}
//
//	globDb.DB().SetMaxIdleConns(config.Config.Mysql.GOMaxIdleConns)
//	globDb.DB().SetMaxOpenConns(config.Config.Mysql.MaxOpenConns)
//	globDb.DB().SetConnMaxLifetime(config.Config.Mysql.ConnMaxLifetime)
//
//	// 打印数据库查询的sql
//	if os.Getenv("DB_DEBUG") == "true" {
//		globDb.LogMode(true)
//	}
//
//	// 创建表 MIGRATE_DB=true
//	if os.Getenv("MIGRATE_DB") == "true" {
//		Migrage()
//	}
//	globDb.SetLogger(logging.GetGormLogger())
//}

func GetDB() (db *gorm.DB) {
	db = globDb
	return
}

func Migrage() {
	logs.Info("begin create table")
	logs.Info("end create table")
}
