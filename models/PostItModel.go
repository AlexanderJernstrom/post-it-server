package models

type PostIt struct {
	ID      uint   `gorm:"primary_key" json:"id"`
	Content string `json:"content"`
	Color   string `json:"color"`
	UserID  uint   `json:"userID"`
	User    User   `json:"user"`
}