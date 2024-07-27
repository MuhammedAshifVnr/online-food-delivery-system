package repository

import (
	"crypto/hmac"
	"crypto/sha256"
	"crypto/subtle"
	"encoding/hex"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/MuhammedAshifVnr/online-food-delivery-system/payment-service/internal/domain/model"
	"github.com/gin-gonic/gin"
	"github.com/razorpay/razorpay-go"
	"gorm.io/gorm"
)

type PaymentRepository interface {
	PaymentConfirmation(c *gin.Context) (string, string)
	NewRazorOrder(orderId string, price uint32) (string, error)
}

type paymentRepository struct {
	db *gorm.DB
}

func NewPaymentRepository(db *gorm.DB) PaymentRepository {
	return &paymentRepository{db: db}
}

func (p *paymentRepository) PaymentConfirmation(c *gin.Context) (string, string) {
	var paymentStore model.PaymentDetails
	var pd = make(map[string]string)
	if err := c.BindJSON(&pd); err != nil {
		log.Fatal("failed to fetch payment data")
	}
	fmt.Println("sig - ", pd["razorpay_signature"], pd["razorpay_order_id"], pd["razorpay_payment_id"])
	err := RazorPaymentVerification(pd["razorpay_signature"], pd["razorpay_order_id"], pd["razorpay_payment_id"])
	if err != nil {
		log.Fatal("-------", err)
	}

	paymentStore.PaymentId = pd["razorpay_payment_id"]
	paymentStore.Status = "success"
	paymentStore.RazorOrderId = pd["razorpay_order_id"]
	fmt.Println(paymentStore)
	p.db.Create(&paymentStore)

	return paymentStore.RazorOrderId, paymentStore.Status
}

func RazorPaymentVerification(sign, orderId, paymentId string) error {
	signature := sign
	secret := os.Getenv("RAZOR_PAY_SECRET")
	data := orderId + "|" + paymentId
	h := hmac.New(sha256.New, []byte(secret))
	_, err := h.Write([]byte(data))
	if err != nil {
		panic(err)
	}
	sha := hex.EncodeToString(h.Sum(nil))
	if subtle.ConstantTimeCompare([]byte(sha), []byte(signature)) != 1 {
		return errors.New("PAYMENT FAILED")
	} else {
		return nil
	}
}

func (p *paymentRepository) NewRazorOrder(orderId string, price uint32) (string, error) {
	client := razorpay.NewClient(os.Getenv("RAZOR_PAY_KEY"), os.Getenv("RAZOR_PAY_SECRET"))
	orderParams := map[string]interface{}{
		"amount":   price * 100,
		"currency": "INR",
		"receipt":  orderId,
	}
	order, err := client.Order.Create(orderParams, nil)
	if err != nil {
		return "", errors.New("PAYMENT NOT INITIATED")
	}

	razorId, _ := order["id"].(string)
	return razorId, nil
}
