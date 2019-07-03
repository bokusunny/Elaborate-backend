package entity

import "time"

type Directory struct {
	ID        int       `json:"id" gorm:"primary_key auto_increment"`
	Name      string    `json:"name" gorm:"not null;default:''"`
	UserID    string    `json:"userId" gorm:"not null;index"`
	CreatedAt time.Time `json:"createdAt" sql:"DEFAULT:current_timestamp"`
	UpdatedAt time.Time `json:"updatedAt" sql:"DEFAULT:current_timestamp on update current_timestamp"`
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
