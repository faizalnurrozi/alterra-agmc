package auth

import (
	"github.com/faizalnurrozi/alterra-agmc/internal/dto"
	"github.com/faizalnurrozi/alterra-agmc/internal/factory"
	res "github.com/faizalnurrozi/alterra-agmc/pkg/util/response"
	"github.com/labstack/echo/v4"
)

type handler struct {
	service Service
}

func NewHandler(f *factory.Factory) *handler {
	return &handler{
		service: NewService(f),
	}
}

func (h *handler) LoginByEmailAndPassword(c echo.Context) error {
	payload := new(dto.ByEmailAndPasswordRequest)
	if err := c.Bind(payload); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err).Send(c)
	}
	if err := c.Validate(payload); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.Validation, err).Send(c)
	}

	user, err := h.service.LoginByEmailAndPassword(c.Request().Context(), payload)
	if err != nil {
		return res.ErrorResponse(err).Send(c)
	}

	return res.SuccessResponse(user).Send(c)
}

func (h *handler) RegisterByEmailAndPassword(c echo.Context) error {
	payload := new(dto.RegisterUserRequestBody)
	if err := c.Bind(payload); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err).Send(c)
	}
	if err := c.Validate(payload); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.Validation, err).Send(c)
	}

	user, err := h.service.RegisterByEmailAndPassword(c.Request().Context(), payload)
	if err != nil {
		return res.ErrorResponse(err).Send(c)
	}

	return res.SuccessResponse(user).Send(c)
}
