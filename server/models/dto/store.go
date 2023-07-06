package dto

import (
	"time"

	"github.com/google/uuid"
)

type StoreRegisterDTOModel struct {
	Id        uuid.UUID
	StoreName string
	Address   string
	NIT       string
	GeoHash   string
	OwnerId   uuid.UUID
	CreatedAt time.Time
}
