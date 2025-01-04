package entities

type User struct {
	Id       string `json:"id" db:"id"`
	Name     string `json:"name" binding:"required"`
	Username string `json:"userName" binding:"required"`
	Password string `json:"password" binding:"required"`
}
