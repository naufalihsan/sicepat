package v1

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/naufalihsan/sicepat/controllers/v1"
)

func SetDestinationRoutes(router *gin.RouterGroup) {
	router.POST("/destination", v1.GetDestinations)
}
