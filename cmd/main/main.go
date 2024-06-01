package main

import (
	"github.com/dendun-nf/gin-golang-web-test/internal"
	"github.com/dendun-nf/gin-golang-web-test/internal/commons"
	"github.com/dendun-nf/gin-golang-web-test/internal/database"
	"github.com/dendun-nf/gin-golang-web-test/internal/todo"
)

func init() {
	commons.GormDb = database.ConnectDatabase()
}

func main() {

	builder := internal.NewApplicationBuilder()
	{
		builder.AddEndpoints(todo.NewEndpoints())
		//builder.AddEndpoints(other.NewEndpoints())
	}

	app := builder.Build()
	app.MapEndpoints()
	app.Run()

	//app := gin.Default()
	//
	//api := app.Group("/api")
	//
	//todos := api.Group("/todos")
	//todos.GET("", service.Index)
	//todos.GET("/:id", service.GetById)
	//todos.POST("", service.Add)
	//todos.PATCH("/:id", service.Update)
	//todos.DELETE("/:id", service.Delete)
	//
	//err := app.Run(":8080")
	//if err != nil {
	//	panic(err)
	//}
}
