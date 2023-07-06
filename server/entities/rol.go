package entities

import "github.com/google/uuid"

type Roles struct {
	Id    uint8          `gorm:"primaryKey;column:id;autoIncrement;not null;unique"`
	Role  string         `gorm:"column:role;type:varchar(50);not null;unique"`
	Users []UserHasRoles `gorm:"foreignKey:role_id"`
}

type UserHasRoles struct {
	UserId   uuid.UUID `gorm:"column:user_id;primaryKey"`
	RoleId   uint8     `gorm:"column:role_id;primaryKey"`
	IsActive bool      `gorm:"column:is_active;type:bool;not null;default:true"`
}
