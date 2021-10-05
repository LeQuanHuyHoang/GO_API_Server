package main

import (
	"GO_API_Server/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/user/get_user/:name", controller.GetUser)
	r.POST("/user/create_user", controller.CreateUser)
	r.PUT("/user/update_user/:name", controller.UpdateUser)
	r.DELETE("/user/delete_user/:name", controller.DeleteUser)
	r.Run()
}
