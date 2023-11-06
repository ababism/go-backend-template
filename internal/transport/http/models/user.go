package models

import (
	"github.com/google/uuid"
	"hw_3/internal/service/domain"
)

// ---- User ----

type UserCreateRequest struct {
	Name string `json:"name" binding:"required"`
}

type UserResponse struct {
	Id      uuid.UUID `json:"id"`
	Name    string    `json:"name"`
	Balance float64   `json:"balance"`
}

func MakeUserResponse(u domain.User) UserResponse {
	return UserResponse{
		Id:      u.Id,
		Name:    u.Name,
		Balance: u.Balance,
	}
}

type TransferRequest struct {
	SenderId   uuid.UUID `json:"sender_id" binding:"required"`
	ReceiverId uuid.UUID `json:"receiver_id" binding:"required"`
	Amount     float64   `json:"amount" binding:"required"`
}

type ClientBalanceResponse struct {
	ClientBalance float64 `json:"client_balance"`
}

type AddBalanceRequest struct {
	ClientId uuid.UUID `json:"receiver_id" binding:"required"`
	Amount   float64   `json:"amount"`
}
