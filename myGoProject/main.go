package main

import (
	"movie-tracker-api/config"
	"movie-tracker-api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()       // Gin router başlatıldı
	config.ConnectDatabase() // Veritabanı bağlantısı sağlanacak
	routes.RegisterRoutes(r) // API route'ları tanımlanacak
	r.Run(":8080")           // Uygulama 8080 portunda başlatılacak
}
