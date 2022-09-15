package controllers

import (
	"day2-crud/config"
	"day2-crud/lib/database"
	"day2-crud/lib/utils"
	"day2-crud/models"
	"day2-crud/routes/requests"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type UserController struct {
	config.HTTPResponse
}

func NewUserController() *UserController {
	return &UserController{}
}

func (uc UserController) GetUsers(ctx echo.Context) (err error) {
	var repositoryUser database.UserRepository
	users, err := repositoryUser.GetAll()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return uc.ResponseOk(ctx, users)
}

func (uc UserController) GetUserByID(ctx echo.Context) (err error) {
	var repositoryUser database.UserRepository
	id, err := strconv.Atoi(ctx.Param("id"))
	fmt.Println(id, "Bener ...")
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	users, err := repositoryUser.GetByID(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return uc.ResponseOk(ctx, users)
}

func (uc UserController) Login(ctx echo.Context) (err error) {
	var repositoryUser database.UserRepository
	req := models.User{}
	err = ctx.Bind(&req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	user, err := repositoryUser.Login(&req)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
	}

	return uc.ResponseOk(ctx, user)
}

func (uc UserController) Create(ctx echo.Context) (err error) {
	var repositoryUser database.UserRepository
	req := new(requests.UserRequest)
	if err = ctx.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err = ctx.Validate(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	hashing := utils.NewHashing()
	password, err := hashing.HashPassword(req.Password)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	userModel := models.User{
		Name:          req.Name,
		Email:         req.Email,
		Password:      password,
		Gender:        req.Gender,
		Nik:           req.Nik,
		BirthDate:     req.BirthDate,
		MarriedStatus: req.MarriedStatus,
		YearOfJoin:    req.YearOfJoin,
	}
	user, err := repositoryUser.Create(userModel)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return uc.ResponseOk(ctx, user)
}

func (uc UserController) Update(ctx echo.Context) (err error) {
	var repositoryUser database.UserRepository
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	req := new(requests.UserRequest)
	if err = ctx.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err = ctx.Validate(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	hashing := utils.NewHashing()
	password, err := hashing.HashPassword(req.Password)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	userModel := models.User{
		Name:          req.Name,
		Email:         req.Email,
		Password:      password,
		Gender:        req.Gender,
		Nik:           req.Nik,
		BirthDate:     req.BirthDate,
		MarriedStatus: req.MarriedStatus,
		YearOfJoin:    req.YearOfJoin,
	}
	user, err := repositoryUser.Update(userModel, id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return uc.ResponseOk(ctx, user)
}

func (uc UserController) Delete(ctx echo.Context) (err error) {
	var repositoryUser database.UserRepository
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	_, err = repositoryUser.Delete(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return uc.ResponseOk(ctx, nil)
}
