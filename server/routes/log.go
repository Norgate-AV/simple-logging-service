package routes

import (
	"net/http"

	"github.com/Norgate-AV/simple-logging-service/models"
	"github.com/gin-gonic/gin"
)

func CreateLog(context *gin.Context) {
	var log *models.Log

	if err := context.ShouldBindJSON(&log); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	if err := log.Save(); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})

		return
	}

	context.JSON(http.StatusCreated, log)
}
