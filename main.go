package main

import (
	"to_do_list/internal/configs"
	"to_do_list/internal/handler"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE"}
	router.Use(cors.New(config))
	handler.RegisterRoutes(router)
	port := configs.LoadPort()
	weburl := configs.LoadWebUrl()
	router.Run(weburl + ":" + port)
}
