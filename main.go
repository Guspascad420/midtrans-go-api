package main

import (
	"midtrans-go-api/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/api/charge", controllers.ChargeRequest)
	r.GET("/api/status/:orderId", controllers.CheckOrder)
	r.Run()
}
