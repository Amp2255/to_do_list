package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"to_do_list/internal/handler"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type APIResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

func SetUpRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.GET("/tasks", handler.ViewAllTasks)
	r.POST("/tasks", handler.SaveTasks)
	r.PUT("/tasks", handler.UpdateATask)
	r.DELETE("/tasks", handler.DeleteTask)
	return r
}

func TestViewTaks(t *testing.T) {
	r := SetUpRouter()

	req, _ := http.NewRequest("GET", "/tasks", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	var resp APIResponse
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	assert.NoError(t, err)
	assert.True(t, resp.Success)

	// Optional: assert the data contains known items
	dataMap, ok := resp.Data.([]interface{})
	assert.True(t, ok)
	assert.Greater(t, len(dataMap), 0)
}
