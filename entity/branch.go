package entity

import "time"

type Branch struct {
	ID             int       `json:"id" gorm:"primary_key auto_increment"`
	Name           string    `json:"name"`
	BaseBranchID   int       `json:"baseBranchId"`
	BaseBranchName string    `json:"baseBranchName"`
	Body           string    `json:"body"`
	State          string    `json:"state"` // TODO: 'open' | 'closed' | 'merged'のEnumチェック
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
}

func NewBranch(name string, baseBranchID int, baseBranchName string, body string, state string) *Branch {
	// TODO: バリデーション
	return &Branch{
		Name:           name,
		BaseBranchID:   baseBranchID,
		BaseBranchName: baseBranchName,
		Body:           body,
		State:          state,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}
}