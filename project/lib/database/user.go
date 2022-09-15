package database

import (
	"day2-crud/config"
	"day2-crud/lib/utils"
	"day2-crud/middlewares"
	"day2-crud/models"
	"errors"
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

func (ur UserRepository) Login(user *models.User) (interface{}, error) {
	var err error
	reqPassword := user.Password
	if err = config.DB.Where("email = ?", user.Email).First(user).Error; err != nil {
		return nil, err
	}

	hashing := utils.NewHashing()
	match := hashing.CheckPasswordHash(reqPassword, user.Password)
	if !match {
		return nil, errors.New("Unauthorized")
	}

	user.Token, err = middlewares.JWTMiddleware{}.CreateToken(int(user.ID))
	if err != nil {
		return nil, err
	}

	if err = config.DB.Save(user).Error; err != nil {
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
	user.Password = data.Password
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
