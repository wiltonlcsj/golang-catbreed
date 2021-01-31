package models

type User struct {
	Id       int64  `json:"id,omitempty"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
