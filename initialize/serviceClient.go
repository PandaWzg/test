package initialize

import (
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	"time"
	"wm-infoflow-api-go/common/db"
	"wm-infoflow-api-go/common/log"
	queue "wm-infoflow-api-go/common/queue/redis"
	"wm-infoflow-api-go/conf"
)

var (
	DefaultDBClient    *gorm.DB
	DefaultRedisClient *redis.Client
	RedisQueue         *queue.Queue
)


type ActiveRecord interface{}

func DBClient() *gorm.DB {
	//if conf.Config.Debug {
	//	_, _ = db.RegisterMySQLDail(conf.Config.SSH)
	//}

	client, err := db.GetMySQL(conf.Config.MySQL["travel"])
	if err != nil {
		log.Fatal("db connection fatal ", err)
	}
	if conf.Config.Debug {
		client.LogMode(true)
	}
	DefaultDBClient = client
	return client
}

func RedisClient() *redis.Client {
	DefaultRedisClient = redis.NewClient(&redis.Options{
		Addr:     viper.GetString("redis.addr"),
		Password: viper.GetString("redis.password"),
		MinIdleConns: 5,
		PoolSize: 200,
	})

	_, err := DefaultRedisClient.Ping().Result()
	if err != nil {
		panic(err)
	}
	return DefaultRedisClient
}

func TxRecovery(tx *gorm.DB) {
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			log.Error("recovery err:", r)
		}
	}()
}

func Lock(key string, expire time.Duration) (locked bool) {
	locked, _ = DefaultRedisClient.SetNX(key, 1, expire).Result()
	return
}

func UnLock(key string) {
	DefaultRedisClient.Del(key).Result()
}