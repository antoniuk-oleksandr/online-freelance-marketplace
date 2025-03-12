package model

import "time"

type User struct {
	ID         int64       `json:"user_id" db:"user_id"`
	About      *string     `json:"about" db:"about"`
	CreatedAt  time.Time   `json:"created_at" db:"created_at"`
	Email      string      `json:"email" db:"email"`
	FirstName  string      `json:"first_name" db:"first_name"`
	Level      float64     `json:"level" db:"level"`
	Password   string      `json:"password" db:"password"`
	PrivateKey *string     `json:"private_key" db:"private_key"`
	PublicKey  *string     `json:"public_key" db:"public_key"`
	Surname    string      `json:"surname" db:"surname"`
	Count      int64       `json:"count" db:"count"`
	Rating     float64     `json:"rating" db:"rating"`
	Username   string      `json:"username" db:"username"`
	Avatar     *string     `json:"avatar" db:"avatar"`
	RoleID     int64       `json:"role_id" db:"role_id"`
	Languages  *[]Language `json:"languages" db:"languages"`
	Skills     *[]Skill    `json:"skills" db:"skills"`
}
