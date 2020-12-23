package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	v1 "github.com/naufalihsan/sicepat/routers/v1"
)

func main() {
	router := gin.Default()
	router.Use(cors.Default())
	api := router.Group("/api")
	{
		v1.SetOriginRoutes(api)
		v1.SetDestinationRoutes(api)
		v1.SetPickupRoutes(api)
		v1.SetTrackingRoutes(api)
	}
	err := router.Run(":8080")
	if err != nil {
		panic(err)
	}
}
