package entity

import "time"

type Commit struct {
	ID        int       `json:"id" gorm:"primary_key auto_increment"`
	Name      string    `json:"name" gorm:"not null;default:''"`
	Body      string    `json:"body" gorm:"not null;default:''"`
	BranchID  int       `json:"branchId" gorm:"not null;index"`
	CreatedAt time.Time `json:"createdAt" sql:"DEFAULT:current_timestamp"`
	UpdatedAt time.Time `json:"updatedAt" sql:"DEFAULT:current_timestamp on update current_timestamp"`
}

func NewCommit(name string, body string, branchID int) *Commit {
	return &Commit{
		Name:      name,
		Body:      body,
		BranchID:  branchID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
