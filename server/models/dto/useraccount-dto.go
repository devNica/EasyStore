package dto

import (
	"time"

	"github.com/google/uuid"
)

type UserRegisterDTOModel struct {
	Id       uuid.UUID
	Email    string
	Password string
	TwoFactorAuth bool
	CreatedAt time.Time
	StatusId uint8
}

type UserLoginDTOModel struct {
	Email    string
	Password string
}
