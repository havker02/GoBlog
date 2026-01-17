package models

import (
  "gorm.io/gorm"
  "github.com/google/uuid"
  "time"
)

type Post struct {
  ID uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
  AuthorId uuid.UUID `gorm:"type:uuid;not null" json:"author_id"`
  Author User `gorm:"foreignKey:AuthorId" json:"author"`
  Title string `gorm:not null" json:"title"`
  Slug string `gorm:"unique;not null" json:"slug"`
  Excerpt string `json:"excerpt"`
  Content string `gorm:"type:text" json:"content"`
  Status string `gorm:"default:draft" json:"status"`
  Visibility string `gorm:"default:public" json:"visibility"`
  PublishedAt time.Time `json:"published_at"`
  Views int64 `gorm:default:0" jso :"views"`
  CommentCount int64 `gorm:"default:0" json:"comment_count"`
  ReactionCount int64 `gorm:"default:0" json:"reaction_count"`
  CreatedAt time.Time `json:"created_at"`
  UpdatedAt time.Time `json:"updated_at"`

  // relations
  Comments []Comment `gorm:"foreignKey:PostId" json:"comments"`
  Reactions []Reaction `gorm:"foreignKey:PostId" json:"reactions"`
  Tags []Tag `gorm:"many2many:post_tags" json:"tags"`
}