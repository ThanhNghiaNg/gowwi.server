package models

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type Transaction struct {
	ID           string    `json:"id"`
	User         string    `json:"user"`
	TypeID       string    `json:"type_id"`
	TypeName     string    `json:"type_name"`
	CategoryID   string    `json:"category"`
	CategoryName string    `json:"category_name"`
	PartnerID    string    `json:"partner_id,omitempty"`
	PartnerName  string    `json:"partner_name,omitempty"`
	Amount       float64   `json:"amount"`
	Date         time.Time `json:"date"`
	Description  string    `json:"description,omitempty"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type CreateTransaction struct {
	User         string    `json:"user"`
	TypeID       string    `json:"type_id"`
	TypeName     string    `json:"type_name"`
	CategoryID   string    `json:"category_id"`
	CategoryName string    `json:"category_name"`
	PartnerID    string    `json:"partner_id"`
	PartnerName  string    `json:"partner_name"`
	Amount       float64   `json:"amount"`
	Date         time.Time `json:"date"`
	Description  string    `json:"description,omitempty"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type UpdateTransaction struct {
	ID          string    `json:"id"`
	User         string    `json:"user"`
	TypeID       string    `json:"type_id"`
	TypeName     string    `json:"type_name"`
	CategoryID   string    `json:"category_id"`
	CategoryName string    `json:"category_name"`
	PartnerID    string    `json:"partner_id"`
	PartnerName  string    `json:"partner_name"`
	Amount       float64   `json:"amount"`
	Date         time.Time `json:"date"`
	Description  string    `json:"description,omitempty"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type CreateTransactionRepository struct {
	User         bson.ObjectID `bson:"user"`
	TypeID       bson.ObjectID `bson:"type_id"`
	TypeName     string        `json:"type_name"`
	CategoryID   bson.ObjectID `bson:"category"`
	CategoryName string        `json:"category_name"`
	PartnerID    bson.ObjectID `bson:"partner_id,omitempty"`
	PartnerName  string        `json:"partner_name,omitempty"`
	Amount       float64       `json:"amount"`
	Date         time.Time     `json:"date"`
	Description  string        `json:"description,omitempty"`
	CreatedAt    time.Time     `json:"created_at"`
	UpdatedAt    time.Time     `json:"updated_at"`
}
