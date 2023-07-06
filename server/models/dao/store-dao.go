package dao

import (
	"time"

	"github.com/google/uuid"
)

type StoreDAOModel struct {
	StoreId   uuid.UUID
	StoreName string
	Address   string
	NIT       string
	GeoHash   string
	IsActive  bool
	CreatedAt time.Time
	OwnerId   uuid.UUID
}
