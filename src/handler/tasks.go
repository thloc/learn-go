package handler

import (
	"demo-todo/src/config"
	"demo-todo/src/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

var task model.Task = model.Task{}

func GetAllTasks(appCtx config.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		tasks, err := task.GetAll(appCtx.GetMainDBConnection())

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, gin.H{"data": tasks})
	}
}
