package entities

import (
	"time"

	"github.com/google/uuid"
)

type Store struct {
	Id        uuid.UUID `gorm:"primaryKey;column:id;type:varchar(36);not null;unique"`
	StoreName string    `gorm:"column:store_name;type:text;size=100;not null;unique"`
	Address   string    `gorm:"column:address;type:text;size=200;not null"`
	CreatedAt time.Time `gorm:"column:created_at;not null"`
	OwnerId   uuid.UUID `gorm:"column:owner_id;primaryKey"`
}
