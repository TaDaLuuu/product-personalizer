package model

import "time"

type User struct {
	UserId    string    `json:"userId,omitempty" db:"user_id,omitempty"`
	FullName  string    `json:"fullName,omitempty" db:"full_name,omitempty"`
	Email     string    `json:"email,omitempty" db:"email,omitempty"`
	Password  string    `json:"-" db:"password,omitempty"`
	Role      string    `json:"role,omitempty" db:"role,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty" db:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty" db:"updated_at,omitempty"`
	Token     string    `json:"token,omitempty" db:"token,omitempty"`
}
