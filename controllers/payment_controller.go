package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
	"midtrans-go-api/models"
	"net/http"
	"os"
	"strconv"
	"time"
)

func ChargeRequest(context *gin.Context) {
	serverKey := os.Getenv("MIDTRANS_SERVER_KEY")
	c := coreapi.Client{}
	c.New(serverKey, midtrans.Sandbox)
	var paymentRequest models.PaymentRequest

	if err := context.ShouldBindJSON(&paymentRequest); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		context.Abort()
		return
	}

	if paymentRequest.PaymentType == "gopay" {
		req := &coreapi.ChargeReqWithMap{
			"payment_type": paymentRequest.PaymentType,
			"transaction_details": map[string]interface{}{
				"order_id":     "MID-GO-TEST-" + Random(),
				"gross_amount": paymentRequest.GrossAmount,
			},
			"gopay": map[string]interface{}{
				"callback_url": "https://midtrans.com/",
			},
		}
		res, _ := c.ChargeTransactionWithMap(req)
		actions := res["actions"].([]interface{})

		context.JSON(http.StatusCreated,
			gin.H{"deeplink_redirect": actions[1].(map[string]interface{})["url"],
				"transaction_id": res["transaction_id"]})

	} else if paymentRequest.PaymentType == "qris" {
		req := &coreapi.ChargeReqWithMap{
			"payment_type": paymentRequest.PaymentType,
			"transaction_details": map[string]interface{}{
				"order_id":     "MID-GO-TEST-" + Random(),
				"gross_amount": paymentRequest.GrossAmount,
			},
		}
		res, _ := c.ChargeTransactionWithMap(req)
		actions := res["actions"].([]interface{})
		context.JSON(http.StatusCreated,
			gin.H{"qr_code_url": actions[0].(map[string]interface{})["url"],
				"transaction_id": res["transaction_id"]})

	} else if paymentRequest.PaymentType == "shopeepay" {
		orderId := "MID-GO-TEST-" + Random()
		req := &coreapi.ChargeReqWithMap{
			"payment_type": paymentRequest.PaymentType,
			"transaction_details": map[string]interface{}{
				"order_id":     orderId,
				"gross_amount": paymentRequest.GrossAmount,
			},
			"shopeepay": map[string]interface{}{"callback_url": "https://midtrans.com/"},
		}
		res, err := c.ChargeTransactionWithMap(req)
		if err != nil {
			panic(err)
		}
		actions := res["actions"].([]interface{})
		context.JSON(http.StatusCreated,
			gin.H{"deeplink_redirect": actions[0].(map[string]interface{})["url"],
				"transaction_id": res["transaction_id"]})
	}
}

func CheckOrder(context *gin.Context) {
	c := coreapi.Client{}
	c.New("SB-Mid-server--_zGSUUz694QDEAzi3Ra9W5z", midtrans.Sandbox)
	orderId := context.Param("orderId")
	res, _ := c.CheckTransaction(orderId)
	context.JSON(http.StatusOK, gin.H{"status": res.TransactionStatus})
}

func Random() string {
	time.Sleep(500 * time.Millisecond)
	return strconv.FormatInt(time.Now().Unix(), 10)
}
