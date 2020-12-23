package v1

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/naufalihsan/sicepat/models/v1"
	"github.com/naufalihsan/sicepat/services"
	"net/http"
)

func CheckTariff(c *gin.Context) {
	var request v1.CheckTariffRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := services.CheckTariff(&request)
	if err != nil {
		c.JSON(http.StatusOK, err)
		return
	}

	c.JSON(http.StatusOK, response)

}

func CalculateInsurance(c *gin.Context) {
	var request v1.CalculateInsuranceRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := services.CalculateInsurance(&request)
	if err != nil {
		c.JSON(http.StatusOK, err)
		return
	}

	c.JSON(http.StatusOK, response)

}

func Pickup(c *gin.Context) {
	var request v1.PickupRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := services.Pickup(&request)
	if err != nil {
		c.JSON(http.StatusOK, err)
		return
	}

	c.JSON(http.StatusOK, response)

}
