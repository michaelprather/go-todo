package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type todo struct {
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

var todos = []todo{
	{Id: 0, Title: "Task 1", Completed: false},
	{Id: 1, Title: "Task 2", Completed: false},
	{Id: 2, Title: "Task 3", Completed: false},
}

// Returns all todos
func getTodos(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, todos)
}

// Creates a todo
func createTodo(c *gin.Context) {
	var newTodo todo

	if err := c.BindJSON(&newTodo); err != nil {
		return
	}

	todos = append(todos, newTodo)

	c.IndentedJSON(http.StatusOK, newTodo)
}

// Updates the todo with the given id.
func updateTodo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	for i, todo := range todos {
		if todo.Id == id {
			todos[i].Completed = !todo.Completed
			break
		}
	}
	c.IndentedJSON(http.StatusOK, todos)
}

// Deletes the todo with the given id.
func deleteTodo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	fmt.Println(id)
	for i, todo := range todos {
		if todo.Id == id {
			todos = append(todos[:i], todos[i+1:]...)
			break
		}
	}
	c.IndentedJSON(http.StatusOK, todos)
}

func main() {
	router := gin.Default()
	router.GET("/todos", getTodos)
	router.POST("/todos", createTodo)
	router.PATCH("/todos/:id", updateTodo)
	router.DELETE("/todos/:id", deleteTodo)
	router.Run(":8080")
}
