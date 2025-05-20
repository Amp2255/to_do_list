package service

import (
	"context"
	"fmt"
	"net/http"
	"time"
	"to_do_list/internal/database"
	"to_do_list/internal/model"
	"to_do_list/internal/utils"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var client = database.Connect()
var tasksCollection *mongo.Collection = database.GetCollection(client, "tasks")

type RequestBody struct {
	Status    string
	Priority  string
	DueBefore time.Time `json:"due_before,omitempty"`
}

func GetAllTasks(ctx *gin.Context) {
	var tasksColl []model.Tasks

	filterVar := bson.D{{}}
	//to find one user only
	results, err := tasksCollection.Find(ctx, filterVar)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.APIresponse{
			Success: false,
			Message: "Error fetching data",
			Error:   err.Error(),
		})
	} else {

		for results.Next(ctx) {
			var task model.Tasks
			if err = results.Decode(&task); err != nil {
				ctx.JSON(http.StatusInternalServerError, utils.APIresponse{
					Success: false,
					Message: "Couldnot find tasks",
					Error:   err.Error(),
				})
			}
			if err == nil {
				tasksColl = append(tasksColl, task)
			}
		}
		ctx.JSON(200, utils.APIresponse{
			Success: true,
			Message: "Task retrieved successfully",
			Data:    tasksColl,
		})

	}

}

func GetAllTasksByFilter(contxt *gin.Context) {

	var tasksColl []model.Tasks
	var reqBody model.Tasks //RequestBody
	if err := contxt.Bind(&reqBody); err != nil {
		contxt.JSON(http.StatusBadRequest, utils.APIresponse{
			Success: false,
			Message: "",
			Data:    tasksColl,
			Error:   err.Error(),
		})
	} else {
		//to find users by filter
		filter := bson.M{}
		if reqBody.Status != "" {
			filter["status"] = reqBody.Status
		}
		if reqBody.Priority != "" {
			filter["priority"] = reqBody.Priority
		}
		// if !reqBody.DueBefore.IsZero() {
		// 	filter["duedate"] = bson.M{"$lt": reqBody.DueBefore}
		// }
		fmt.Println("requestFilter is :", filter["status"])
		results, err := tasksCollection.Find(context.TODO(), filter)

		if err != nil {
			contxt.JSON(http.StatusInternalServerError, utils.APIresponse{
				Success: false,
				Message: "Error fetching data",
				Error:   err.Error(),
			})
		}
		defer results.Close(context.TODO())
		for results.Next(contxt) {
			var task model.Tasks
			err = results.Decode(&task)
			// err != nil {
			// 	ctx.JSON(http.StatusInternalServerError, utils.APIresponse{
			// 		Success: false,
			// 		Message: "Couldnot find tasks",
			// 		Error:   err.Error(),
			// 	})
			// }

			if err == nil {

				tasksColl = append(tasksColl, task)
			}
		}
		contxt.JSON(http.StatusOK, utils.APIresponse{
			Success: true,
			Message: "Task retrieved successfully",
			Data:    tasksColl,
		})

	}
}

func CreateTasks(contxt *gin.Context) {
	var anewTask model.Tasks
	if err := contxt.BindJSON(&anewTask); err != nil {
		contxt.JSON(http.StatusBadRequest, utils.APIresponse{
			Success: false,
			Message: "Error binding json",
			Data:    anewTask,
			Error:   err.Error(),
		})
		return
	}

	anewTask.Id = primitive.NewObjectID()
	results, err := tasksCollection.InsertOne(context.TODO(), anewTask)
	if err != nil {
		contxt.JSON(http.StatusInternalServerError, utils.APIresponse{
			Success: false,
			Message: "Error saving new task",
			Data:    anewTask,
			Error:   err.Error(),
		})
		return
	} else {
		contxt.JSON(http.StatusOK, utils.APIresponse{
			Success: true,
			Message: "Task saved successfully",
			Data:    results,
		})
	}
}

func UpdateATask(contxt *gin.Context) { ///tasks/:id
	fmt.Println("********** UpdateATask", contxt.Query("id"))
	taskIdToUpdate := contxt.Query("id")
	objId, err := primitive.ObjectIDFromHex(taskIdToUpdate)
	fmt.Println("********** Object Id", objId)
	if err != nil {
		contxt.JSON(http.StatusBadRequest, utils.APIresponse{
			Success: false,
			Message: "Invalid task ID",
			Data:    objId,
			Error:   err.Error(),
		})
		return
	}

	var reqBody model.Tasks
	if err := contxt.Bind(&reqBody); err != nil {
		contxt.JSON(http.StatusBadRequest, utils.APIresponse{
			Success: false,
			Message: "",
			Data:    reqBody,
			Error:   err.Error(),
		})
	} else {
		reqBody.UpdatedAt = time.Now()
		fmt.Println("reqBody to update :", reqBody)
		dataToUpdate := bson.M{
			"$set": bson.M{
				"title":       reqBody.Title,
				"description": reqBody.Description,
				"status":      reqBody.Status,
				"priority":    reqBody.Priority,
				"due_date":    reqBody.DueDate,
				"updated_at":  reqBody.UpdatedAt,
			},
		}

		results, err := tasksCollection.UpdateByID(context.TODO(), objId, dataToUpdate)
		if err != nil {
			contxt.JSON(http.StatusInternalServerError, utils.APIresponse{
				Success: false,
				Message: "Failed to Update Task",
				Error:   err.Error(),
			})
			return
		}

		if results.MatchedCount == 0 {
			contxt.JSON(http.StatusNotFound, utils.APIresponse{
				Success: false,
				Message: "Task not found",
				Error:   "",
			})
			return
		}

		contxt.JSON(http.StatusOK, utils.APIresponse{
			Success: true,
			Message: "Task updated successfully",
			Data:    results.MatchedCount,
		})

	}
}

func DeleteTask(contxt *gin.Context) {
	taskIdToDelete := contxt.Query("id")
	objId, err := primitive.ObjectIDFromHex(taskIdToDelete)
	fmt.Println("********** Object Id to delete", objId)
	if err != nil {
		contxt.JSON(http.StatusBadRequest, utils.APIresponse{
			Success: false,
			Message: "Invalid task ID",
			Data:    objId,
			Error:   err.Error(),
		})
		return
	}
	results, err := tasksCollection.DeleteOne(context.TODO(), bson.M{"_id": objId})
	if err != nil {
		contxt.JSON(http.StatusInternalServerError, utils.APIresponse{
			Success: false,
			Message: "Failed to delete Task",
			Error:   err.Error(),
		})
		return
	}
	if results.DeletedCount == 0 {
		contxt.JSON(http.StatusNotFound, utils.APIresponse{
			Success: false,
			Message: "Task not found",
			Error:   "",
		})
		return
	}

	contxt.JSON(http.StatusOK, utils.APIresponse{
		Success: true,
		Message: "Task deleted successfully",
		Data:    results.DeletedCount,
	})
}
