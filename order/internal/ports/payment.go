package ports

import (
	"context"

	"github.com/avi1152111680093/microservices/order/internal/application/core/domain"
)

type PaymentPort interface {
	Charge(context.Context, *domain.Order) error
}
