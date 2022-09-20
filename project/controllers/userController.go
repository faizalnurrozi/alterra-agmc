package controllers

import (
	"github.com/faizalnurrozi/alterra-agmc/config"
	"github.com/faizalnurrozi/alterra-agmc/lib/database"
	"github.com/faizalnurrozi/alterra-agmc/lib/utils"
	"github.com/faizalnurrozi/alterra-agmc/models"
	"github.com/faizalnurrozi/alterra-agmc/routes/requests"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type UserController struct {
	config.Contract
}

func NewUserController(contract config.Contract) *UserController {
	return &UserController{Contract: contract}
}

func (uc UserController) GetUsers(ctx echo.Context) (err error) {
	if ctx.Request().Method == http.MethodGet && ctx.Request().Header.Get("Authorization") != "" {
		var repositoryUser database.UserRepository
		users, err := repositoryUser.GetAll()
		if err != nil {
			return uc.ResponseOk(ctx, http.StatusBadRequest, nil)
		}
		return uc.ResponseOk(ctx, http.StatusOK, users)
	} else if ctx.Request().Header.Get("Authorization") == "" {
		return uc.ResponseOk(ctx, http.StatusUnauthorized, nil)
	}
	return uc.ResponseOk(ctx, http.StatusMethodNotAllowed, nil)
}

func (uc UserController) GetUserByID(ctx echo.Context) (err error) {
	if ctx.Request().Method == http.MethodGet && ctx.Request().Header.Get("Authorization") != "" {
		var repositoryUser database.UserRepository
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			return uc.ResponseOk(ctx, http.StatusBadRequest, nil)
		}

		users, err := repositoryUser.GetByID(id)
		if err != nil {
			return uc.ResponseOk(ctx, http.StatusBadRequest, nil)
		}

		return uc.ResponseOk(ctx, http.StatusOK, users)
	} else if ctx.Request().Header.Get("Authorization") == "" {
		return uc.ResponseOk(ctx, http.StatusUnauthorized, nil)
	}
	return uc.ResponseOk(ctx, http.StatusMethodNotAllowed, nil)
}

func (uc UserController) Login(ctx echo.Context) (err error) {
	if ctx.Request().Method == http.MethodPost {
		var repositoryUser database.UserRepository
		req := requests.LoginRequest{}
		err = ctx.Bind(&req)
		if err != nil {
			return uc.ResponseOk(ctx, http.StatusBadRequest, nil)
		}

		if err = ctx.Validate(req); err != nil {
			return uc.ResponseOk(ctx, http.StatusBadRequest, nil)
		}

		var userReqLogin = models.User{
			Email:    req.Email,
			Password: req.Password,
		}
		user, err := repositoryUser.Login(&userReqLogin)
		if err != nil {
			return uc.ResponseOk(ctx, http.StatusUnauthorized, nil)
		}

		ctx.Set("user", user)

		return uc.ResponseOk(ctx, http.StatusOK, user)
	}
	return uc.ResponseOk(ctx, http.StatusMethodNotAllowed, nil)
}

func (uc UserController) Create(ctx echo.Context) (err error) {
	if ctx.Request().Method == http.MethodPost {
		var repositoryUser database.UserRepository
		req := models.User{}
		err = ctx.Bind(&req)
		if err != nil {
			return uc.ResponseOk(ctx, http.StatusBadRequest, nil)
		}

		if err = ctx.Validate(req); err != nil {
			return uc.ResponseOk(ctx, http.StatusBadRequest, nil)
		}

		hashing := utils.NewHashing()
		password, err := hashing.HashPassword(req.Password)
		if err != nil {
			return uc.ResponseOk(ctx, http.StatusBadRequest, nil)
		}

		req.Password = password
		user, err := repositoryUser.Create(&req)
		if err != nil {
			return uc.ResponseOk(ctx, http.StatusBadRequest, nil)
		}

		return uc.ResponseOk(ctx, http.StatusCreated, user)
	}
	return uc.ResponseOk(ctx, http.StatusMethodNotAllowed, nil)
}

func (uc UserController) Update(ctx echo.Context) (err error) {
	if ctx.Request().Method == http.MethodPut && ctx.Request().Header.Get("Authorization") != "" {
		var repositoryUser database.UserRepository
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			return uc.ResponseOk(ctx, http.StatusBadRequest, nil)
		}

		req := models.User{}
		err = ctx.Bind(&req)
		if err != nil {
			return uc.ResponseOk(ctx, http.StatusBadRequest, nil)
		}

		if err = ctx.Validate(req); err != nil {
			return uc.ResponseOk(ctx, http.StatusBadRequest, nil)
		}

		hashing := utils.NewHashing()
		password, err := hashing.HashPassword(req.Password)
		if err != nil {
			return uc.ResponseOk(ctx, http.StatusBadRequest, nil)
		}

		req.Password = password
		user, err := repositoryUser.Update(&req, id)
		if err != nil {
			return uc.ResponseOk(ctx, http.StatusBadRequest, nil)
		}

		return uc.ResponseOk(ctx, http.StatusOK, user)
	} else if ctx.Request().Header.Get("Authorization") == "" {
		return uc.ResponseOk(ctx, http.StatusUnauthorized, nil)
	}
	return uc.ResponseOk(ctx, http.StatusMethodNotAllowed, nil)
}

func (uc UserController) Delete(ctx echo.Context) (err error) {
	if ctx.Request().Method == http.MethodDelete && ctx.Request().Header.Get("Authorization") != "" {
		var repositoryUser database.UserRepository
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			return uc.ResponseOk(ctx, http.StatusBadRequest, nil)
		}

		_, err = repositoryUser.Delete(id)
		if err != nil {
			return uc.ResponseOk(ctx, http.StatusMethodNotAllowed, nil)
		}

		return uc.ResponseOk(ctx, http.StatusOK, nil)
	} else if ctx.Request().Header.Get("Authorization") == "" {
		return uc.ResponseOk(ctx, http.StatusUnauthorized, nil)
	}
	return uc.ResponseOk(ctx, http.StatusMethodNotAllowed, nil)
}
