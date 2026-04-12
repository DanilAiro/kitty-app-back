package models

type User struct {
	Email    string `json:"email" binding:"required,email" bson:"email"`
	Password string `json:"password" binding:"required,min=8" bson:"password"`
}