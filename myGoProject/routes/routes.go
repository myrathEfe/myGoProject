package routes

import (
	"movie-tracker-api/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	r.GET("/movies", controllers.GetMovies)
	r.POST("/movies", controllers.CreateMovie)
}
