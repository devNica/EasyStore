package dao

import (
	"time"

	"github.com/google/uuid"
)

type RolDAOModel struct {
	Id  uint8
	Rol string
}

type FindUserDAOModel struct {
	Id        uuid.UUID
	Email     string
	Password  string
	CreatedAt time.Time
}
