package model

import "golang.org/x/crypto/bcrypt"

type User struct {
	ID       uint   `gorm:"primary_key"`
	Username string `gorm:"unique;not null"`
	Email    string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
}

func (u *User) HashPassword() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}
