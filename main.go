package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type todo struct {
	Id        string `json:"id"`
	Item      string `json:"item"`
	Completed bool   `json:"completed"`
}

var todos = []todo{
	{Id: "1", Item: "Learned Golang", Completed: false},
	{Id: "2", Item: "Learned PyTorch", Completed: false},
	{Id: "3", Item: "Learned Python", Completed: false},
}

func getToDos(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, todos)
}

func getToDoById(id string) (*todo, error) {
	for i, t := range todos {
		if t.Id == id {
			return &todos[i], nil
		}
	}

	return nil, errors.New("todo not found")
}

func getToDo(context *gin.Context) {
	id := context.Param("id")

	todo, err := getToDoById(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
		return
	}

	context.IndentedJSON(http.StatusOK, todo)
}

func addToDo(context *gin.Context) {
	var newToDo todo

	if err := context.BindJSON(&newToDo); err != nil {
		return
	}

	todos = append(todos, newToDo)

	context.IndentedJSON(http.StatusCreated, newToDo)
}

func main() {
	router := gin.Default()
	router.GET("/todos", getToDos)
	router.POST("/todos", addToDo)
	router.GET("/todos/:id", getToDo)
	router.Run("localhost:9090")
}
