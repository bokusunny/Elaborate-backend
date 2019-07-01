package entity

import "time"

type Directory struct {
	ID        int       `json:"id" gorm:"primary_key auto_increment"`
	Name      string    `json:"name"`
	UserID    string    `json:"userId"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Branches  []Branch
}

func NewDirectory(name string, userID string) *Directory {
	return &Directory{
		Name:      name,
		UserID:    userID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
