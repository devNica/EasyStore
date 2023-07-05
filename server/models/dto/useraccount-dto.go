package dto

import (
	"time"

	"github.com/google/uuid"
)

type UserRegisterDTOModel struct {
	Id            uuid.UUID
	Email         string
	Password      string
	TwoFactorAuth bool
	CreatedAt     time.Time
	StatusId      uint8
}

type UserLoginDTOModel struct {
	Email    string
	Password string
}

type PersonalInfoDTOModel struct {
	PhoneNumber   string
	DNI           string
	FirstName     string `gorm:"column:firstname"`
	LastName      string `gorm:"column:lastname"`
	Address       string
	BirthDate     string `gorm:"column:birthdate"`
	TwoFactorAuth bool
	UpdatedAt     time.Time
	StatusId      uint8
}
