package entity

import (
	"demo-todo/src/config"
	"demo-todo/src/helper"
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

		c.JSON(http.StatusOK, helper.BuildResponse(tasks))
	}
}

func CreateTask(appCtx config.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var tasks model.Task

		if err := c.ShouldBind(&tasks); err != nil {
			c.JSON(401, gin.H{
				"error": err.Error(),
			})

			return
		}

		if err := task.CreateTask(appCtx.GetMainDBConnection(), &tasks); err != nil {
			c.JSON(401, gin.H{
				"error": err.Error(),
			})

			return
		}

		c.JSON(http.StatusOK, helper.BuildResponse(tasks))
	}
}
