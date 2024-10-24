package main

import (
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

func main() {
	router := gin.Default()
	router.GET("/todos", getToDos)
	router.Run("localhost:9090")
}
