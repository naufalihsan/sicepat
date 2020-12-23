package v1

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/naufalihsan/sicepat/controllers/v1"
)

func SetPickupRoutes(router *gin.RouterGroup) {
	router.POST("/check-tariff", v1.CheckTariff)
	router.POST("/calculate-insurance", v1.CalculateInsurance)
	router.POST("/pickup", v1.Pickup)
}
