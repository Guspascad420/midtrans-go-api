package main

import (
	"github.com/gin-gonic/gin"
	"midtrans-go-api/controllers"
)

func main() {
	r := gin.Default()

	r.POST("/api/charge", controllers.ChargeRequest)
	r.POST("/api/otp/generate", controllers.GenerateAndSendOTP)
	r.POST("/api/otp/validate", controllers.ValidateOTP)
	r.GET("/api/status/:orderId", controllers.CheckOrder)
	r.Run()
}
