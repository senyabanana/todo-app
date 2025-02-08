package handler

import (
	"github.com/senyabanana/todo-app/internal/entity"
	
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type errorResponse struct {
	Message string `json:"message"`
}

type statusResponse struct {
	Status string `json:"status"`
}

func newErrorResponse(c *gin.Context, statusCode int, message string) {
	logrus.Error(message)
	c.AbortWithStatusJSON(statusCode, errorResponse{Message: message})
}

type getAllListsResponse struct {
	Data []entity.TodoList `json:"data"`
}
