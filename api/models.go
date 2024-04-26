package api

import "time"

type User struct {
	ID        string `gorm:"primaryKey;autoIncrement"`
	Username  string `gorm:"unique"`
	Password  string
	CreatedAt time.Time
}

func CreateNewUser(username, password string) (*User, error) {
	user := &User{
		Username:  username,
		Password:  password, // Hash password here before saving
		CreatedAt: time.Now(),
	}
	result := db.Create(user)
	return user, result.Error
}
