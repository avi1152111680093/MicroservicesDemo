package grpc

import (
	"context"
	"fmt"
	"log"

	"github.com/avi1152111680093/microservices/payment/internal/application/core/domain"
	payment "github.com/avi1152111680093/microservices/payment/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (a Adapter) Create(ctx context.Context, request *payment.CreatePaymentRequest) (*payment.CreatePaymentResponse, error) {
	newPayment := domain.NewPayment(request.UserId, request.OrderId, request.TotalPrice)

	// This functions calls the DB for saving payment information and returns the payment information created
	log.Printf("Charging for the Order")
	result, err := a.api.Charge(ctx, newPayment)
	if err != nil {
		return nil, status.New(codes.Internal, fmt.Sprintf("failed to charge. %v ", err)).Err()
	}
	return &payment.CreatePaymentResponse{PaymentId: result.ID}, nil
}