package models

import (
  "gorm.io/gorm"
  "github.com/google/uuid"
  "time"
)

type Tag struct {
  ID uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
  Name steing `gorm:"unique;not null" json:"name"`
  Slug string `gorm:"unique;not null" json:"slug"`
  CreatedAt time.Time `json:"created_at"`
  UpdateddAt time.Time `json:"updated_at"`
}