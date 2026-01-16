package models

import (
  "gorm.io/gorm"
  "github.com/google/uuid"
  "time"
)

type Reaction struct {
  ID uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
  PostId uuid.UUID `gorm:"type:uuid;not null" json:"post_id"`
  UserId uuid.UUID `gorm:"type:uuid;not null" json:"user_id"`
  Type string `gorm:"not null;default:like" json:"type"`
  CreatedAt time.Time `json:"created_at"`
  UpdatedAt time.Time `json:"updated_at"`
}