package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/xlzd/gotp"
	"math/rand"
	"midtrans-go-api/models"
	"net/http"
	"net/smtp"
	"os"
)

var secret = os.Getenv("OTP_SECRET")
var hotp = gotp.NewDefaultHOTP(secret)
var generatedOTP string

func GenerateAndSendOTP(context *gin.Context) {
	var emailRequest models.EmailRequest
	if err := context.ShouldBindJSON(&emailRequest); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	generatedOTP = hotp.At(rand.Int())
	auth := smtp.PlainAuth("", "paacad@gmail.com",
		os.Getenv("EMAIL_HOST_PASSWORD"), "smtp.gmail.com")
	to := []string{emailRequest.Email}
	msg := []byte("To: " + emailRequest.Email + "\r\n" +
		"Here's your OTP code\r\n" +
		"\r\n" + generatedOTP + "\r\n")
	err := smtp.SendMail("smtp.gmail.com:587", auth, "paacad@gmail.com", to, msg)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "email sent", "otp": generatedOTP})
}

func ValidateOTP(context *gin.Context) {
	var otpRequest models.OTPRequest

	if err := context.ShouldBindJSON(&otpRequest); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if otpRequest.Otp != generatedOTP {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "OTP is invalid"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "OTP is valid"})
}
