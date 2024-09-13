package task

import (
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
	"108HW/internal/websocket"
)

func CreateTaskHandler(hub *websocket.Hub) gin.HandlerFunc {
	return func(c *gin.Context) {
		var task Task
		if err := c.ShouldBindJSON(&task); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
			return
		}
		createdTask := CreateTask(task)
		hub.Broadcast <- []byte("Task created: " + createdTask.Title)
		c.JSON(http.StatusCreated, createdTask)
	}
}

func GetTasksHandler(c *gin.Context) {
	tasks := GetTasks()
	c.JSON(http.StatusOK, tasks)
}

func UpdateTaskHandler(hub *websocket.Hub) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		var task Task
		if err := c.ShouldBindJSON(&task); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
			return
		}
		updatedTask, ok := UpdateTask(id, task)
		if !ok {
			c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
			return
		}
		hub.Broadcast <- []byte("Task updated: " + updatedTask.Title)
		c.JSON(http.StatusOK, updatedTask)
	}
}

func DeleteTaskHandler(hub *websocket.Hub) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		if ok := DeleteTask(id); !ok {
			c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
			return
		}
		hub.Broadcast <- []byte("Task deleted with ID: " + strconv.Itoa(id))
		c.Status(http.StatusNoContent)
	}
}
