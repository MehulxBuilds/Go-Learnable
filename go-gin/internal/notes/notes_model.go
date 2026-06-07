package notes

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Note struct {
	ID primitive.ObjectID `bson:"_id"  json:"id"`

	Title string `bson:"title" json:"title"`

	Content string `bson:"content" json:"content"`

	Read bool `bson:"read" json:"read"`

	CreatedAt time.Time `bson:"createdAt" json:"createdAt"`

	UpdatedAt time.Time `bson:"updatedAt" json:"updatedAt"`
}

type CreateNoteRequest struct {
	Title string `json:"title" binding:"required"`

	Content string `json:"content" binding:"required"`

	Read bool `json:"read"`
}

type UpdateNoteRequest struct {
	Title string `json:"title" binding:"required"`

	Content string `json:"content" binding:"required"`

	Read bool `json:"read"`
}
