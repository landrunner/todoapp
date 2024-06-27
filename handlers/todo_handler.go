package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/landrunner/todo/models"
)

var todos = []models.Todo{
	{ID: 1, Title: "Learn go", Status: "assinged"},
	{ID: 2, Title: "Make Tod", Status: "assinged"},
}

func ShowIndexPage(c *gin.Context) {
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
	newTodo.ID = uint(len(todos) + 1)
	todos = append(todos, newTodo)
	c.Redirect(http.StatusFound, "/")
}

func GetTodos(c *gin.Context) {
	c.JSON(http.StatusOK, todos)
}
