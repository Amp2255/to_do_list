package handler

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {

	router.GET("/tasks", ViewAllTasks)
	//router.POST("/tasks", ViewTasksByFilter)
	router.POST("/tasks", SaveTasks)
	router.PUT("/tasks", UpdateATask)
	router.DELETE("/tasks", DeleteTask)
}
