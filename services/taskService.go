package services

import (
	"goprj/entities"
	"goprj/infrastructures"
	"log"
)

func GetTaskOfUser(email string) (tasks []entities.Task, err error) {
	db := infrastructures.GetDB()

	log.Println(email)

	db.Where("user_email = ?", email).Find(&tasks)

	if db.Error != nil {
		log.Println("ERROR: can not find user tasks")
		return nil, db.Error
	}

	return tasks, nil
}

func AddNewTask(email string, title string) (*entities.Task, error) {
	db := infrastructures.GetDB()

	task := entities.Task{
		UserEmail: email,
		Title:     title,
	}
	db.Create(&task)

	if db.Error != nil {
		log.Println("ERROR: can not add new task")
		return nil, db.Error
	}

	return &task, nil
}
