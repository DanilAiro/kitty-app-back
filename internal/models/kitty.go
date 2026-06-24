package models

type Kitty struct {
	Id     string `json:"id" binding:"required" bson:"id"`
	URL    string `json:"url" binding:"required" bson:"url"`
}
