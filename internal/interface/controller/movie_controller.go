package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type (
	MovieController struct {}

	GetMovieRequest struct {
	Name string `json:"name" binding:"required"`
	}
)

func (mc *MovieController) getMovie(c *gin.Context) {
	var req GetMovieRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro, payload inválido": err.Error()})
		return
	}
}