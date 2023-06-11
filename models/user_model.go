package models

import (
	"log"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// to ignore empty fields and make the field required, respectively.
// type User struct {
// 	Id       primitive.ObjectID `json:"id,omitempty"`
// 	Name     string             `json:"name,omitempty" validate:"required"`
// 	Location string             `json:"location,omitempty" validate:"required"`
// 	Password string             `json:"password"`
// 	Username string `json:"username" gorm:"unique"`
// 	Email    string `json:"email" gorm:"unique"`
// }

type UserLogin struct {
	Id       primitive.ObjectID `json:"id,omitempty"`
	Password string             `json:"password" validate:"required"`
	Email    string `json:"email" gorm:"unique"`
}

type User struct {
	gorm.Model
	Id       primitive.ObjectID `json:"id,omitempty"`
	Name     string             `json:"name,omitempty" validate:"required"`
	Location string             `json:"location,omitempty" validate:"required"`
	Password string             `json:"password"`
	Username string `json:"username" gorm:"unique"`
	Email    string `json:"email" gorm:"unique"`
}

func (user *User) HashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}
	user.Password = string(bytes)
	return nil
}
func (user *User) CheckPassword(providedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(providedPassword))
	if err != nil {
		return err
	}
	return nil
}

func (user *UserLogin) CheckLoginPassword(providedPassword string) error {
	log.Print("Dungji")
	log.Print(providedPassword)
	log.Print("Dungji")
	log.Print(user.Password)
	log.Print("Dungji")
	err := bcrypt.CompareHashAndPassword([]byte(providedPassword), []byte(user.Password))
	if err != nil {
		return err
	}
	return nil
}
