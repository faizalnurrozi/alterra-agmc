package controllers

import (
	"day2-crud/config"
	"day2-crud/lib/database"
	"day2-crud/lib/utils"
	"day2-crud/models"
	"day2-crud/routes/requests"
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

	return uc.ResponseOk(ctx, http.StatusOK, users)
}

func (uc UserController) GetUserByID(ctx echo.Context) (err error) {
	var repositoryUser database.UserRepository
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	users, err := repositoryUser.GetByID(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return uc.ResponseOk(ctx, http.StatusOK, users)
}

func (uc UserController) Login(ctx echo.Context) (err error) {
	var repositoryUser database.UserRepository
	req := requests.LoginRequest{}
	err = ctx.Bind(&req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err = ctx.Validate(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	var userReqLogin = models.User{
		Email:    req.Email,
		Password: req.Password,
	}
	user, err := repositoryUser.Login(&userReqLogin)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
	}

	ctx.Set("user", user)

	return uc.ResponseOk(ctx, http.StatusOK, user)
}

func (uc UserController) Create(ctx echo.Context) (err error) {
	var repositoryUser database.UserRepository
	req := models.User{}
	err = ctx.Bind(&req)
	if err != nil {
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

	req.Password = password
	user, err := repositoryUser.Create(&req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return uc.ResponseOk(ctx, http.StatusOK, user)
}

func (uc UserController) Update(ctx echo.Context) (err error) {
	var repositoryUser database.UserRepository
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	req := models.User{}
	err = ctx.Bind(&req)
	if err != nil {
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

	req.Password = password
	user, err := repositoryUser.Update(&req, id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return uc.ResponseOk(ctx, http.StatusOK, user)
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

	return uc.ResponseOk(ctx, http.StatusOK, nil)
}
