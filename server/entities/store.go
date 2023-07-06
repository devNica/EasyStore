package entities

import (
	"time"

	"github.com/google/uuid"
)

type Store struct {
	Id        uuid.UUID `gorm:"primaryKey;column:id;type:varchar(36);not null;unique"`
	StoreName string    `gorm:"column:store_name;type:text;size=100;not null;unique"`
	Address   string    `gorm:"column:address;type:text;size=200;not null"`
	NIT       string    `gorm:"column:nit;type:varchar(100);null"`
	GeoHash   string    `gorm:"column:geo_hash;type:varchar(20);null"`
	IsActive  bool      `gorm:"column:is_active;type:bool;not null;default=true"`
	OwnerId   uuid.UUID `gorm:"column:owner_id;primaryKey"`
	CreatedAt time.Time `gorm:"column:created_at;not null"`
	UpdatedAt time.Time `gorm:"column:updated_at;null"`
}
