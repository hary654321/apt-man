package redis

import (
	"github.com/go-redis/redis"
	"go.uber.org/zap"
	"zrDispatch/common/log"
	"zrDispatch/core/config"
	"zrDispatch/core/slog"
)

var client *redis.Client

func GetClient() *redis.Client {
	if client == nil {
		slog.Println(slog.WARN, "初始化redis客户端")
		client = redis.NewClient(&redis.Options{
			Addr:     config.CoreConf.Server.Redis.Addr,
			Password: config.CoreConf.Server.Redis.PassWord,
		})

		err := client.Ping().Err()
		if err != nil {
			log.Error("connect redis failed", zap.String("addr", config.CoreConf.Server.Redis.Addr))
			return nil
		}

	}
	return client
}
