package grpc

import (
	"context"
	"github.com/avi1152111680093/microservices/order/internal/application/core/domain"
	order "github.com/avi1152111680093/microservices/order/proto"
)

func (a Adapter) Create(ctx context.Context, request *order.CreateOrderRequest) (*order.CreateOrderResponse, error) {
	var orderItems []domain.OrderItem
	for _, orderItem := range request.OrderItems {
		orderItems = append(orderItems, domain.OrderItem{
			ProductCode: orderItem.ProductCode,
			UnitPrice: orderItem.UnitPrice,
			Quantity: orderItem.Quantity,
		})
	}

	newOrder := domain.NewOrder(request.UserId, orderItems)
	
	// This function calls the DB for saving the data, and then calls the Payment service
	result, err := a.api.PlaceOrder(ctx, newOrder)
	if err != nil {
		return nil, err
	}
	
	return &order.CreateOrderResponse{OrderId: result.ID}, nil
}