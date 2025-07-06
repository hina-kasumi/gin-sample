package main

import (
	"goprj/entities"
	"goprj/infrastructures"
	"goprj/middlewares"
	"goprj/routes"
	"goprj/services"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		return
	} // khai báo để đọc từ .env

	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	host := os.Getenv("DB_HOST")

	db := infrastructures.OpenDbConnection(username, password, dbName, host)
	if db != nil {
		db.AutoMigrate(&entities.User{}, &entities.Task{})
	}

	services.InitRedis(os.Getenv("REDIS_ADDRESS"))

	// 2. Tạo router
	r := gin.Default()
	r.Use(middlewares.Cors())
	r.Use(middlewares.UserLoaderMiddleware())
	// 3. Đăng ký các routes
	routes.RegisterUserRoutes(r)
	routes.RegisterTaskRoutes(r)
	routes.RegisterAuthRoutes(r)

	// 4. Chạy server
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Không thể khởi động server:", err)
	}
}
