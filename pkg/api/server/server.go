package server

import (
	"github.com/gin-gonic/gin"
	"github.com/minhajuddin/todo/pkg/api/handlers"
)

func Start(addr string) {
	r := gin.Default()
	WireUpRoutes(r)
	r.Run(addr)
}

func WireUpRoutes(r *gin.Engine) {
	r.GET("/tasks", handlers.ListTasks)
}
