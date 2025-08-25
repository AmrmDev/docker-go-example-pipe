package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"docker-go-example-pipe/internal/usecase"
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

	movie, err := mc.usecase.GetMovieByName(req.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro ao buscar filme": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"filme": movie})
}