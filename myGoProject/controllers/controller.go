package controllers

import (
	"movie-tracker-api/config"
	"movie-tracker-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetMovies(c *gin.Context) {
	var movies []models.Movie
	config.DB.Find(&movies)
	c.JSON(http.StatusOK, movies)
}

func CreateMovie(c *gin.Context) {
	var movie models.Movie
	if err := c.ShouldBindJSON(&movie); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Create(&movie)
	c.JSON(http.StatusCreated, movie)
}
