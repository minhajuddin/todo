package handlers

import "github.com/gin-gonic/gin"

func ListTasks(c *gin.Context) {
	c.Writer.Write([]byte("Tasks"))
}
