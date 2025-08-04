package models

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type UserLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserRegister struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type User struct {
	ID        string    `json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	FullName  string    `json:"full_name"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	Address   string    `json:"address"`
	IsAdmin   bool      `json:"is_admin"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserRepository struct {
	ID        bson.ObjectID `bson:"_id"`
	Username  string        `json:"username"`
	Password  string        `json:"password"`
	FullName  string        `json:"full_name"`
	Email     string        `json:"email"`
	Phone     string        `json:"phone"`
	Address   string        `json:"address"`
	IsAdmin   bool          `json:"is_admin"`
	CreatedAt time.Time     `json:"created_at"`
	UpdatedAt time.Time     `json:"updated_at"`
}

// define UpdateUser from User struct
type UpdateUser struct {
	ID        string    `json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	FullName  string    `json:"full_name"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	Address   string    `json:"address"`
	IsAdmin   bool      `json:"is_admin"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreateUser struct {
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	Email     string    `json:"email"`
	IsAdmin   bool      `json:"is_admin"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
