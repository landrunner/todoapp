package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/landrunner/todo/models"
)

func ShowIndexPage(c *gin.Context) {
	val, _ := c.Get("db")
	db, ok := val.(models.DataSource)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "DB Connection is broken"})
	}

	todos, err := db.FetchTodos()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "DB Connection is broken"})
	}
	c.HTML(http.StatusOK, "index.html", todos)
}

func ShowCreatePage(c *gin.Context) {
	c.HTML(http.StatusOK, "create.html", nil)
}

func CreateTodoHTML(c *gin.Context) {
	var newTodo models.Todo
	if err := c.ShouldBind(&newTodo); err != nil {
		c.HTML(http.StatusBadRequest, "create.html", gin.H{"error": err.Error()})
		return
	}

	val, _ := c.Get("db")
	db, ok := val.(models.DataSource)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "DB Connection is broken"})
	}

	_ = db.AddTodo(newTodo)
	c.Redirect(http.StatusFound, "/")
}
