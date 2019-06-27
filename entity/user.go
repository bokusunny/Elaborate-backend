package entity

import "time"

// User has user data sent by firebase cloud functions
type User struct {
	ID        int       `json:"id" gorm:"primary_key auto_increment"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// NewUser returns user data with given parameters
func NewUser(name string, email string) *User {
	return &User{
		Name:      name,
		Email:     email,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
