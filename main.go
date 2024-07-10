package main

import (
	"github.com/gin-gonic/gin"
	"github.com/landrunner/todo/handlers"
	"github.com/landrunner/todo/models"
)

func main() {
	var db = models.DataSource{}
	db.InitDB("testdb.sqlite3")
	r := gin.Default()
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})
	r.Static("/static", "./static")
	r.LoadHTMLGlob("templates/*")
	r.GET("/", handlers.ShowIndexPage)
	r.GET("/create", handlers.ShowCreatePage)
	r.POST("/create", handlers.CreateTodoHTML)
	r.Run()
}
