package main

import (
	"goprj/entities"
	"goprj/infrastructures"
	"goprj/routes"
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

	// 2. Tạo router
	r := gin.Default()
	// 3. Đăng ký các routes
	routes.RegisterUserRoutes(r)
	routes.RegisterTaskRoutes(r)

	// 4. Chạy server
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Không thể khởi động server:", err)
	}
}
