package initalize

import (
	"sync"

	"github.com/spf13/viper"
)

var once sync.Once

func init() {
	// 应用生命周期 只执行一次 链接数据库操作
	once.Do(func() {
		if viper.GetBool("Mysql.Enable") {
			InitDatabaseMysql()
		}
		if viper.GetBool("Redis.Enable") {
			initDatabaseRedis()
		}
		if viper.GetBool("Zap.Enable") {
			InitDatabaseLogger()
		}
	})
}
