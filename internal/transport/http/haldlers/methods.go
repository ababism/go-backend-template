package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"hw_3/internal/service/domain"
	"hw_3/internal/transport/http/models"
	"net/http"
)

func bindRequestBody(c *gin.Context, obj any) bool {
	if err := c.BindJSON(obj); err != nil {
		newErrorResponse(c, http.StatusBadRequest, domain.ErrIncorrectBody.Error())
		return false
	}
	return true
}

func newErrorResponse(c *gin.Context, statusCode int, message string) {
	logrus.Infof("%s: %d %s", c.Request.URL, statusCode, message)
	c.AbortWithStatusJSON(statusCode, models.Error{Message: message})
}
func errorResponse(c *gin.Context, statusCode int, err error) {
	logrus.Infof("%s: %d %s", c.Request.URL, statusCode, err.Error())
	c.AbortWithStatusJSON(statusCode, models.Error{Message: err.Error()})
}

func mapErrorToCode(c *gin.Context, err error) int {
	switch err {
	case domain.ErrInternal:
		return http.StatusInternalServerError
	case domain.ErrNotFound:
		return http.StatusNotFound
	case domain.ErrIncorrectBody:
		return http.StatusBadRequest
	case domain.ErrTokenInvalid:
		return http.StatusUnauthorized
	case domain.ErrAccessDenied:
		return http.StatusUnauthorized
	case domain.ErrAlreadyExists:
		return http.StatusBadRequest
	case domain.ErrBadUUID:
		return http.StatusForbidden
	default:
		return http.StatusInternalServerError
	}
}

// parseUUIDFromParam makes Error response if it couldn't parse token, returns true if everything is ok
func parseUUIDFromParam(c *gin.Context) (uuid.UUID, bool) {
	id := c.Param("id")
	itemUUID, err := uuid.Parse(id)
	if err != nil {
		newErrorResponse(c, mapErrorToCode(c, domain.ErrBadUUID), domain.ErrBadUUID.Error())
		return uuid.Nil, false
	}
	return itemUUID, true
}

//func parseUUID(c *gin.Context, id string) (uuid.UUID, bool) {
//	itemUUID, err := uuid.Parse(id)
//	if err != nil {
//		newErrorResponse(c, mapErrorToCode(c, domain.ErrBadUUID), domain.ErrBadUUID.Error())
//		return uuid.Nil, false
//	}
//	return itemUUID, true
//}
