package compoments

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"self-wechat/config"
	// 导入postgres驱动
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	pgdns = fmt.Sprintf(
		"host=%s port=%d user=%s dbname=%s sslmode=disable password=%s ",
		config.Config.Pgsql.Host,
		config.Config.Pgsql.Port,
		config.Config.Pgsql.Username,
		config.Config.Pgsql.Dbname,
		config.Config.Pgsql.Password,
	)
	globlDbPg *gorm.DB
)

func init() {
	//var (
	//	err error
	//)
	//if globlDbPg, err = gorm.Open("postgres", pgdns); err != nil {
	//	log.Fatal(err)
	//}
	//// 连接池
	//globlDbPg.DB().SetMaxIdleConns(config.Config.Pgsql.MaxIdleConns)
	//globlDbPg.DB().SetMaxOpenConns(config.Config.Pgsql.MaxOpenConns)
	//globlDbPg.DB().SetConnMaxLifetime(config.Config.Pgsql.ConnMaxLifetime)
	////需要设置环境变量 DB_DEBUG=true
	//if os.Getenv("DB_DEBUG") == "true" {
	//	globlDbPg.LogMode(true)
	//}
	//////需要设置环境变量 MIGRATE_DB=true
	//if os.Getenv("MIGRATE_DB") == "true" {
	//	MigragePG()
	//}
	//globlDbPg.SetLogger(logging.GetGormLogger())
}
func MigragePG() {
	logs.Info("begin create table")
	logs.Info("end create table")
}

// GetPGDB ...
func GetPGDB() (db *gorm.DB) {
	db = globlDbPg
	return
}
