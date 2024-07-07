package models

import (
	"chatter-backend-go/src/configs"
	"time"
)

type User struct {
	ID            string     `gorm:"primaryKey;type:uuid;default:gen_random_uuid();not null" json:"id"`
	Fullname      string     `gorm:"not null" json:"fullname"`
	Username      string     `gorm:"not null;unique" json:"username"`
	Email         string     `gorm:"not null;unique" json:"email"`
	Password      string     `gorm:"not null" json:"password"`
	Image         *string    `json:"image,omitempty"`
	PhoneNumber   *string    `json:"phone_number,omitempty"`
	EmailVerified bool       `gorm:"default:false" json:"email_verified"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at"`
	DeletedAt     *time.Time `sql:"index" json:"deleted_at,omitempty"`
}

func SelectAllUsers() []*User {
	var users []*User
	configs.DB.Limit(200).Find(&users)
	return users
}

func SelectUser(id string) *User {
	var user User
	configs.DB.First(&user, "id = ?", id)
	return &user
}

func CreateUser(userData *User) error {
	result := configs.DB.Create(&userData)
	return result.Error
}

func UpdateUser(userData *User) error {
	result := configs.DB.Model(&User{}).Where("id = ?", userData.ID).Updates(userData)
	return result.Error
}

func DeleteProfilePicture(id string) error {
	result := configs.DB.Model(&User{}).Where("id = ?", id).Update("image", "")
	return result.Error
}

func DeleteUser(id string) error {
	result := configs.DB.Delete(&User{}, "id = ?", id)
	return result.Error
}

func VerifyEmail(email string) error {
	result := configs.DB.Model(&User{}).Where("email = ?", email).Update("email_verified", true)
	return result.Error
}

func FindUsername(username string) *User {
	var user User
	configs.DB.First(&user, "username = ?", username)
	return &user
}

func UpdateUserPassword(email string, password string) error {
	result := configs.DB.Model(&User{}).Where("email = ?", email).Update("password", password)
	return result.Error
}
