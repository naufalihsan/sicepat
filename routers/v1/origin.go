package v1

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/naufalihsan/sicepat/controllers/v1"
)

func SetOriginRoutes(router *gin.RouterGroup) {
	router.POST("/province", v1.GetProvinces)
	router.POST("/city", v1.GetCities)
	router.POST("/district", v1.GetDistricts)
	router.POST("/subdistrict", v1.GetSubDistricts)
}
