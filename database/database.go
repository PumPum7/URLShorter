package database

import (
	"context"
	"github.com/go-redis/redis"
	"urlShortener/utils"
)

var rdb *redis.Client
var ctx = context.Background()

func init() {

	rdb = redis.NewClient(&redis.Options{
		Addr:     utils.Cfg.Database.Domain,
		Password: utils.Cfg.Database.Password,
		Username: utils.Cfg.Database.Username,
		DB:       utils.Cfg.Database.DBName,
	})
}

func GetItem(key string) string {
	result, _ := rdb.HGet(ctx, "urls", key).Result()
	utils.Log.Info(result)
	return result
}

func SetItem(key string, value string) {
	result, _ := rdb.HSet(ctx, "urls", key, value).Result()
	utils.Log.Info(result)
}

func CheckKey(key string) bool {
	result, err := rdb.HExists(ctx, "url", key).Result()
	if err != nil {
		return false
	} else {
		return result
	}
}
