package dao

import (
	"time"

	"github.com/google/uuid"
)

type RolDAOModel struct {
	Id   uint8
	Role string
}

type FindUserDAOModel struct {
	Id        uuid.UUID
	Email     string
	Password  string
	CreatedAt time.Time
}

type AccounStatusDAOModel struct {
	Status string
}
