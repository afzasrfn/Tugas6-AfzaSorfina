package services

import (
	"context"
	"fmt"
	"time"

	cm "Tugas6-AfzaSorfina/git/customer/common"
)

//SubscriberServices is service definition
type PaymentServices interface {
	OrderHandler(context.Context, cm.Message) cm.Message
	CustomerHandler(context.Context, cm.Customers) cm.Customers
	ProductHandler(context.Context, cm.Products) cm.Products
	FastPayHandler(context.Context, cm.FastPayRequest) cm.FastPayResponse
	CallHandler(context.Context, cm.FastPayRequest) cm.FastPayResponse
	TripsHandler(context.Context, cm.MyTripsRequest) cm.MyTripsResponse
}

type PaymentService struct{}

type ServiceMiddleware func(PaymentServices) PaymentServices


func utc() string {
	return time.Now().Format("2006-01-02 15:04:05 +0700")
}

func panicRecovery() {
	if r := recover(); r != nil {
		fmt.Printf("Recovering from panic: %v \n", r)
	}
}