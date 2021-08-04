package main

import (
	"demo-todo/src/config"
	"demo-todo/src/entity"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db := config.ConnectDB()

	if err := runService(db); err != nil {
		log.Fatalln(err)
	}
}

func runService(db *gorm.DB) error {
	appCtx := config.NewAppContext(db)
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// CRUD
	v1 := r.Group("/v1")

	restaurants := v1.Group("/tasks")
	{
		restaurants.GET("", entity.GetAllTasks(appCtx))
	}

	return r.Run()
}
