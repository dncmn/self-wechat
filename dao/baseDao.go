package dao

import (
	"github.com/jinzhu/gorm"
	"self-wechat/compoments"
	"self-wechat/utils/logging"
)

var (
	logger      = logging.GetLogger()
	redisClient *compoments.RedisInstance
	db          *gorm.DB
	pgdb        *gorm.DB
)

func init() {
	db = compoments.GetDB()
	redisClient = compoments.GetRedisClient()
	pgdb = compoments.GetPGDB()
}
