package entities

import (
	"time"

	"github.com/google/uuid"
)

type KYCReviewRequest struct {
	Id             uuid.UUID `gorm:"primaryKey;column:id;type:varchar(36);not null;unique"`
	UserRef        string    `gorm:"column:user_ref;type:varchar(36);not null"`
	AdminRef       string    `gorm:"column:admin_ref;type:varchar(36);not null"`
	PreRevStatus   string    `gorm:"column:pre_rev_status;type:varchar(100);not null"`
	PostRevStatus  string    `gorm:"column:post_rev_status;type:varchar(100);null"`
	CreatedAt      time.Time `gorm:"column:created_at;not null"`
	UpdatedAt      time.Time `gorm:"column:updated_at;null"`
	IsAssigned     bool      `gorm:"column:is_assigned;type:bool;not null;default:false"`
	ReviewStatusId uint8     `gorm:"column:review_status_id;primaryKey"`
}

type ReviewStatus struct {
	Id      uint8              `gorm:"primaryKey;column:id;autoIncrement;not null;unique"`
	Status  string             `gorm:"column:status;type:varchar(100);not null;unique"`
	Reviews []KYCReviewRequest `gorm:"foreignKey:review_status_id"`
}
