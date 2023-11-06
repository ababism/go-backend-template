package domain

import (
	"context"
	"github.com/google/uuid"
)

type UserRepository interface {
	Create(ctx context.Context, user User) (User, error)
	Get(ctx context.Context, id uuid.UUID) (User, error)
	GetBalance(ctx context.Context, userID uuid.UUID) (float64, error)
	AddToBalance(ctx context.Context, userID uuid.UUID, amount float64) (float64, error)
	TransferCurrency(ctx context.Context, senderID uuid.UUID, receiverID uuid.UUID, amount float64) (newSenderBalance float64, err error)
}
