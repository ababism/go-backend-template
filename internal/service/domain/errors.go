package domain

import "errors"

var (
	ErrTokenInvalid  = errors.New("token invalid")
	ErrIncorrectBody = errors.New("incorrect json body")
	ErrInternal      = errors.New("server internal error")
	ErrNotFound      = errors.New("not found")
	ErrAccessDenied  = errors.New("access denied")
	ErrAlreadyExists = errors.New("element already exists")
	ErrBadUUID       = errors.New("could not get correct uuid")
)

func FilterErrors(err error) error {
	switch err {
	case ErrTokenInvalid, ErrIncorrectBody, ErrNotFound, ErrAccessDenied, ErrAlreadyExists, ErrBadUUID:
		return err
	default:
		return ErrInternal
	}
}
