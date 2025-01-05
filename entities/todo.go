package entities

type TodoList struct {
	Id          string `json:"id" db:"id"`
	Title       string `json:"title" binding:"required" db:"title"`
	Description string `json:"description" db:"description"`
}

type UserList struct {
	Id     string
	UserId string
	ListId string
}

type TodoItem struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}

type ListItem struct {
	Id     string
	ListId string
	UserId string
}
