package v1

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/naufalihsan/sicepat/controllers/v1"
)

func SetTrackingRoutes(router *gin.RouterGroup) {
	router.POST("/tracking", v1.Tracking)
}
