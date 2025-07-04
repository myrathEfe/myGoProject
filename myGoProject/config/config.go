package config

import (
	"fmt"
	"log"

	"movie-tracker-api/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

//gorm.DB, GORM ORM (Object Relational Mapper) tarafından sağlanan ve veritabanıyla iletişimi yöneten ana nesnedir.

func ConnectDatabase() {
	dsn := "host=localhost user=postgres password=1234 dbname=movie_tracker port=5432 sslmode=disable"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Veritabanına bağlanılamadı...", err)
	}

	err = database.AutoMigrate(&models.Movie{})
	if err != nil {
		log.Fatal("Tablo oluşturulamadı...:", err)
	}

	DB = database
	fmt.Println("Veritabanına bağlantı başarılı")
}
