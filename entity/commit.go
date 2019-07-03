package entity

import "time"

type Commit struct {
	ID        int       `json:"id" gorm:"primary_key auto_increment"`
	Name      string    `json:"name"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func NewCommit(name string, body string) *Commit {
	return &Commit{
		Name:      name,
		Body:      body,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
