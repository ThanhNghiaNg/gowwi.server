package models

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type Category struct {
	ID          string    `json:"id"`
	User        string    `json:"user"`
	Name        string    `json:"name"`
	Description string    `json:"description,omitempty"`
	Type        string    `json:"type"`
	UsedTime    int       `json:"used_time"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type CreateCategory struct {
	User        string    `json:"user"`
	Name        string    `json:"name"`
	Description string    `json:"description,omitempty"`
	Type        string    `json:"type"`
	UsedTime    int       `json:"used_time"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type UpdateCategory struct {
	ID          string    `json:"id"`
	User        string    `json:"user"`
	Name        string    `json:"name"`
	Description string    `json:"description,omitempty"`
	Type        string    `json:"type"`
	UsedTime    int       `json:"used_time"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type CreateCategoryRepository struct {
	User        bson.ObjectID `bson:"user"`
	Name        string        `json:"name"`
	Description string        `json:"description,omitempty"`
	Type        bson.ObjectID `bson:"type"`
	UsedTime    int           `json:"used_time"`
	CreatedAt   time.Time     `json:"created_at"`
	UpdatedAt   time.Time     `json:"updated_at"`
}

type CategoryRepository struct {
	ID          bson.ObjectID `bson:"_id"`
	User        bson.ObjectID `bson:"user"`
	Name        string        `json:"name"`
	Description string        `json:"description,omitempty"`
	Type        bson.ObjectID `bson:"type"`
	UsedTime    int           `json:"used_time"`
	CreatedAt   time.Time     `json:"created_at"`
	UpdatedAt   time.Time     `json:"updated_at"`
}
