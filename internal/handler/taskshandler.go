package handler

import (
	"fmt"
	"to_do_list/internal/service"

	"github.com/gin-gonic/gin"
)

func ViewAllTasks(contxt *gin.Context) {
	service.GetAllTasks(contxt)

}
func ViewTasksByFilter(ctx *gin.Context) {
	fmt.Println("inside ViewTasksByFilter **************8")
	service.GetAllTasksByFilter(ctx)
}

func SaveTasks(ctx *gin.Context) {
	service.CreateTasks(ctx)
}

func UpdateATask(ctx *gin.Context) {
	fmt.Println("********** UpdateATask", ctx.Query("editingId"))
	service.UpdateATask(ctx)
}

func DeleteTask(ctx *gin.Context) {
	service.DeleteTask(ctx)
}
