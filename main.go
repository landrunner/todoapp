package main

import (
	"github.com/gin-gonic/gin"
	"github.com/landrunner/todo/handlers"
)

func main() {
	r := gin.Default()
	r.Static("/static", "./static")
	r.LoadHTMLGlob("templates/*")

	r.GET("/", handlers.ShowIndexPage)
	r.GET("/create", handlers.ShowCreatePage)
	r.POST("/create", handlers.CreateTodoHTML)
	r.GET("/todo", handlers.GetTodos)
	r.Run()
}
