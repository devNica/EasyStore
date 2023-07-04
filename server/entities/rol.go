package entities

import "github.com/google/uuid"

type Rol struct {
	Id    uint8         `gorm:"primaryKey;column:id;autoIncrement;not null;unique"`
	Rol   string        `gorm:"column:rol;type:varchar(50);not null;unique"`
	Users []UserHasRole `gorm:"foreignKey:rol_id"`
}

type UserHasRole struct {
	UserId   uuid.UUID `gorm:"column:user_id;primaryKey"`
	RolId    uint8     `gorm:"column:rol_id;primaryKey"`
	IsActive bool      `gorm:"column:is_active;type:bool;not null;default:true"`
}
