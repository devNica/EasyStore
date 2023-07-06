package dto

import (
	"time"

	"github.com/google/uuid"
)

type UserRelationShipWithStoreDTO struct {
	OwnerId string
	StoreId string
}

type StoreRegisterDTOModel struct {
	Id        uuid.UUID
	StoreName string
	Address   string
	NIT       string
	GeoHash   string
	OwnerId   uuid.UUID
	CreatedAt time.Time
}

type UpdateStoreDTOModel struct {
	StoreName string
	Address   string
	NIT       string
	GeoHash   string
	UpdatedAt time.Time
}
