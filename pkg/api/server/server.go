package server

import (
	"github.com/gin-gonic/gin"
	"github.com/minhajuddin/todo/pkg/api/handlers"
)

func Start(addr string) {
	r := gin.Default()

	r.GET("/tasks", handlers.ListTasks)

	r.Run(addr)
}
