package entities

import (
	"time"

	"github.com/google/uuid"
)

type File struct {
	Filename  uuid.UUID        `gorm:"primaryKey;column:filename;type:varchar(36);not null;unique"`
	Filetype  string           `gorm:"column:filetype;type:varchar(10);not null"`
	Filesize  int              `gorm:"column:filesize;type:int4;not null"`
	Binary    []byte           `gorm:"column:binary;type:bytea;not nul"`
	CreatedAt time.Time        `gorm:"column:created_at;not null"`
	UpdatedAt time.Time        `gorm:"column:updated_at;null"`
	Stores    []StoreHasAssets `gorm:"foreignKey:file_id;references:filename"`
}

type StoreHasAssets struct {
	StoreId uuid.UUID `gorm:"column:store_id;primaryKey"`
	FileId  uuid.UUID `gorm:"column:file_id;primaryKey"`
}
