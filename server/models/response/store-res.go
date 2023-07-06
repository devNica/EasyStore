package response

import (
	"time"

	"github.com/google/uuid"
)

type StoreResponseModel struct {
	StoreId   uuid.UUID `json:"storeId"`
	StoreName string    `json:"storeName"`
	Address   string    `json:"address"`
	NIT       string    `json:"nit"`
	Latitude  float64   `json:"latitude"`
	Longitude float64   `json:"longitude"`
	IsActive  bool      `json:"isActive"`
	CreatedAt time.Time `json:"createdAt"`
}
