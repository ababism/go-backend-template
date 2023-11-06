package handlers

import (
	"github.com/gin-gonic/gin"
	"hw_3/internal/service/domain"
	"hw_3/internal/transport/http/models"
	"net/http"
)

type UserHandler struct {
	s domain.UserService
}

func NewUserHandler(s domain.UserService) *UserHandler {
	return &UserHandler{s: s}
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	req := models.UserCreateRequest{}
	if !bindRequestBody(c, &req) {
		return
	}

	res, err := h.s.Create(c, domain.User{
		Name: req.Name,
	})

	if err != nil {
		newErrorResponse(c, mapErrorToCode(c, err), err.Error())
		return
	}
	c.JSON(http.StatusOK, models.UserResponse{Id: res.Id, Name: res.Name, Balance: res.Balance})
}

func (h *UserHandler) GetUser(c *gin.Context) {
	// parse from "/user/:id"
	id, ok := parseUUIDFromParam(c)
	if !ok {
		return
	}
	user, err := h.s.Get(c, id)
	if err != nil {
		newErrorResponse(c, mapErrorToCode(c, err), err.Error())
		return
	}
	c.JSON(http.StatusOK, models.MakeUserResponse(user))
}

func (h *UserHandler) TransferBalance(c *gin.Context) {
	// If here will be possibility we should take senderId from JWT
	req := models.TransferRequest{}
	if !bindRequestBody(c, &req) {
		return
	}
	senderBalance, err := h.s.TransferCurrency(c, req.SenderId, req.ReceiverId, req.Amount)

	if err != nil {
		newErrorResponse(c, mapErrorToCode(c, err), err.Error())
		return
	}
	c.JSON(http.StatusOK, models.ClientBalanceResponse{ClientBalance: senderBalance})
}

func (h *UserHandler) AddToUserBalance(c *gin.Context) {
	// If here will be possibility we should take clientId from other service
	req := models.AddBalanceRequest{}
	if !bindRequestBody(c, &req) {
		return
	}
	updatedBalance, err := h.s.AddToBalance(c, req.ClientId, req.Amount)

	if err != nil {
		newErrorResponse(c, mapErrorToCode(c, err), err.Error())
		return
	}
	c.JSON(http.StatusOK, models.ClientBalanceResponse{ClientBalance: updatedBalance})
}
