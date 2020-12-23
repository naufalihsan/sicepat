package v1

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/naufalihsan/sicepat/models/v1"
	"github.com/naufalihsan/sicepat/services"
	"net/http"
)

func Tracking(c *gin.Context) {
	var request v1.TrackingRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := services.Tracking(&request)
	if err != nil {
		c.JSON(http.StatusOK, err)
		return
	}

	c.JSON(http.StatusOK, response)

}
