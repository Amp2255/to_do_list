package handler

import (
	"to_do_list/internal/service"

	"github.com/gin-gonic/gin"
)

func ViewAllTasks(contxt *gin.Context) {
	service.GetAllTasks(contxt)

}
func ViewTasksByFilter(ctx *gin.Context) {
	service.GetAllTasksByFilter(ctx)
}

func SaveTasks(ctx *gin.Context) {
	service.CreateTasks(ctx)
}

func UpdateATask(ctx *gin.Context) {
	service.UpdateATask(ctx)
}

func DeleteTask(ctx *gin.Context) {
	service.DeleteTask(ctx)
}
