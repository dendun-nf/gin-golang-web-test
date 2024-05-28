package main

import (
	_ "github.com/dendun-nf/gin-golang-web-test/database"
	"github.com/dendun-nf/gin-golang-web-test/internal/todo/service"
	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()

	api := app.Group("/api")

	todos := api.Group("/todos")
	todos.GET("", service.Index)
	todos.GET("/:id", service.GetById)
	todos.POST("", service.Add)
	todos.PATCH("/:id", service.Update)
	todos.DELETE("/:id", service.Delete)

	err := app.Run(":8080")
	if err != nil {
		panic(err)
	}
}
