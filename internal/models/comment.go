package models

import (
  "gorm.io/gorm"
  "github.com/google/uuid"
  "time"
)

type Comment struct {
  ID uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
  PostId uuid.UUID `gorm:"type:uuid" json:"post_id"`
  UserId uuid.UUID `gorm:"type:uuid" json:"user_id"`
  ParentId uuid.UUID `gorm:"type:uuid" json:"parent_id"`
  Content string `gorm:"not null" json:"content"`
  IsApproved bool `gorm:"default:true" json:"is_approved"`
  CreatedAt time.Time `json:"created_at"`
  UpdatedAt time.Time `json:"updated_at"`

  // relations
  User []User `gorm:"foreignKey:UserId" json:"user"`
  Post []Post `gorm:"foreignKey:PostId" json:"post"`
  Parent []Comment `gorm:"foreignKey:ParentId" json:"parent"`
}
