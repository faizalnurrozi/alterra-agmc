package database

import (
	"day2-crud/config"
	"day2-crud/models"
)

type UserRepository struct {
}

func (ur UserRepository) GetAll() (interface{}, error) {
	var users []models.User

	if err := config.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (ur UserRepository) GetByID(id int) (interface{}, error) {
	var user models.User

	if err := config.DB.First(&user, id).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (ur UserRepository) Create(user models.User) (interface{}, error) {
	if err := config.DB.Create(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (ur UserRepository) Update(data models.User, id int) (interface{}, error) {
	var user models.User

	if err := config.DB.Find(&user, id).Error; err != nil {
		return nil, err
	}

	user.Name = data.Name
	user.Email = data.Email
	user.Gender = data.Gender
	user.Nik = data.Nik
	user.BirthDate = data.BirthDate
	user.MarriedStatus = data.MarriedStatus
	user.YearOfJoin = data.YearOfJoin
	if err := config.DB.Save(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (ur UserRepository) Delete(id int) (interface{}, error) {
	var user models.User

	if err := config.DB.Delete(&user, id).Error; err != nil {
		return nil, err
	}

	return user, nil
}
