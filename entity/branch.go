package entity

import "time"

type Branch struct {
	ID             int       `json:"id" gorm:"primary_key auto_increment"`
	Name           string    `json:"name" gorm:"not null;default:''"`
	DirectoryID    int       `json:"directoryId" gorm:"not null;index"`
	BaseBranchID   int       `json:"baseBranchId"`
	BaseBranchName string    `json:"baseBranchName"`
	Body           string    `json:"body" gorm:"not null;default:''"`
	State          string    `json:"state" gorm:"not null;default:'open'"` // TODO: 'open' | 'closed' | 'merged'のEnumチェック
	CreatedAt      time.Time `json:"createdAt" sql:"DEFAULT:current_timestamp"`
	UpdatedAt      time.Time `json:"updatedAt" sql:"DEFAULT:current_timestamp on update current_timestamp"`
	Commits        []Commit
}

func NewBranch(name string, directoryID int, baseBranchID int, baseBranchName string, body string, state string) *Branch {
	// TODO: バリデーション
	return &Branch{
		Name:           name,
		DirectoryID:    directoryID,
		BaseBranchID:   baseBranchID,
		BaseBranchName: baseBranchName,
		Body:           body,
		State:          state,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}
}
