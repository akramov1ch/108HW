package main

import (
	"log"
	"108HW/internal/task"
	"108HW/internal/websocket"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	hub := websocket.NewHub()
	go hub.Run()

	r.POST("/task", task.CreateTaskHandler(hub))
	r.GET("/tasks", task.GetTasksHandler)
	r.PUT("/tasks/:id", task.UpdateTaskHandler(hub))
	r.DELETE("/tasks/:id", task.DeleteTaskHandler(hub))

	r.GET("/ws", func(c *gin.Context) {
		websocket.ServeWs(hub, c.Writer, c.Request)
	})

	log.Println("Server running on :8080")
	log.Fatal(r.Run(":8080"))
}
