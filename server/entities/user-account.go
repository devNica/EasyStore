package entities

import (
	"time"

	"github.com/google/uuid"
)

type UserAccount struct {
	Id            uuid.UUID      `gorm:"primaryKey;column:id;type:varchar(36);unique"`
	Email         string         `gorm:"colum:email;type:varchar(255);not null;unique"`
	Password      string         `gorm:"column:password;type:varchar(255);not null"`
	PhoneNumber   string         `gorm:"column:phone_number;type:varchar(30);null"`
	FirstName     string         `gorm:"column:firstname;type:text;size:100;null"`
	LastName      string         `gorm:"column:lastname;type:text;size:100;null"`
	Address       string         `gorm:"column:address;type:text;size:200;null"`
	BirthDate     string         `gorm:"column:birthdate;type:varchar(10);null"`
	DNI           string         `gorm:"column:dni;type:varchar(50);null"`
	TwoFactorAuth bool           `gorm:"column:two_factor_auth;type:bool;default:false"`
	IsRoot        bool           `gorm:"column:is_root;type:bool;not null;default:false"`
	CreatedAt     time.Time      `gorm:"column:created_at"`
	UpdatedAt     time.Time      `gorm:"column:updated_at"`
	StatusId      uint8          `gorm:"column:status_id;primaryKey"`
	Roles         []UserHasRoles `gorm:"foreignKey:user_id"`
	Stores        []Store        `gorm:"foreignKey:owner_id"`
}

type AccountStatus struct {
	Id        uint8         `gorm:"primaryKey;column:id;autoIncrement;not null;unique"`
	Status    string        `gorm:"column:status;type:varchar(50);not null;unique"`
	AccStatus []UserAccount `gorm:"foreignKey:status_id"`
}
