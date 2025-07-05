package services

import (
	"errors"
	"goprj/entities"
	"goprj/infrastructures"
	"log"
)

func FindAllUser() (users []entities.User, err error) {
	result := infrastructures.GetDB().Find(&users)

	if result.Error != nil {
		log.Println("ERROR: ", result.Error)
		err = result.Error
	}
	return
}

func FindOneUser(condition entities.User) (*entities.User, error) {
	db := infrastructures.GetDB().First(&condition)

	if db.Error != nil {
		return nil, db.Error
	}

	return &condition, nil

}

func NewUser(user *entities.User) error {
	existing := entities.User{Email: user.Email}
	if u, _ := FindOneUser(existing); u != nil {
		log.Println("ERROR: Email is existing!")
		return errors.New("email is existing")
	}

	err := user.SetPassword(user.Password)
	if err != nil {
		return errors.New("can not bcrypt password")
	}

	db := infrastructures.GetDB()
	result := db.Create(&user)

	if result.Error != nil {
		log.Println("ERROR: ", result.Error)
	}

	return result.Error
}
