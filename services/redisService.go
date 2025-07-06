package services

import (
	"context"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

var redisClient *redis.Client
var ctx = context.Background()

func InitRedis(addr string) {
	redisClient = redis.NewClient(&redis.Options{
		Addr: addr,
		DB:   0,
	})

	// Kiểm tra kết nối
	_, err := redisClient.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Không thể kết nối Redis: %v", err)
	}

	log.Println("Kết nối Redis thành công")
}

func SetRedisValue(key string, value interface{}) error {
	err := redisClient.Set(ctx, key, value, 0).Err()
	if err != nil {
		log.Println("Can not set value: ", err)
		return err
	}

	return nil
}

func SetRedisExpire(key string, value interface{}, exp int64) error {
	// Set key trước nếu chưa có
	duration := time.Duration(exp-time.Now().Unix()) * time.Second
	err := redisClient.Set(ctx, key, value, duration).Err()
	if err != nil {
		log.Println("Can not set key:", err)
		return err
	}

	// Thiết lập thời gian hết hạn tại thời điểm cụ thể
	log.Println("Sẽ hết hạn sau:", duration)

	return nil
}

func GetRedisValue(key string) (interface{}, error) {
	val, err := redisClient.Get(ctx, key).Result()

	if err == redis.Nil {
		log.Println("Key không tồn tại")
	} else if err != nil {
		log.Println("Lỗi khi lấy dữ liệu:", err)
	} else {
		log.Println("Giá trị của key là:", val)
	}

	return val, err
}
