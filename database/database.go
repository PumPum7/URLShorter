package database

import (
	"github.com/go-redis/redis"
	"urlShortener/utils"
)

var rdb *redis.Client

func init() {

	rdb = redis.NewClient(&redis.Options{
		Addr:     utils.Cfg.Database.Domain,
		Password: utils.Cfg.Database.Password,
		DB:       utils.Cfg.Database.DBName,
	})
}

func GetItem(key string) string {
	result, _ := rdb.HGet("urls", key).Result()
	utils.Log.Info(result)
	return result
}

func SetItem(key string, value string) {
	result, _ := rdb.HSet("urls", key, value).Result()
	utils.Log.Info(result)
}

func CheckKey(key string) bool {
	result, err := rdb.HExists("url", key).Result()
	if err != nil {
		return false
	} else {
		return result
	}
}
