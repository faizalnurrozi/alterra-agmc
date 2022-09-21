package user

import (
	"github.com/faizalnurrozi/alterra-agmc/internal/dto"
	"github.com/faizalnurrozi/alterra-agmc/internal/factory"
	"github.com/faizalnurrozi/alterra-agmc/internal/pkg/util"
	pkgdto "github.com/faizalnurrozi/alterra-agmc/pkg/dto"
	res "github.com/faizalnurrozi/alterra-agmc/pkg/util/response"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

type handler struct {
	service Service
}

func NewHandler(f *factory.Factory) *handler {
	return &handler{
		service: NewService(f),
	}
}

func (h *handler) Get(c echo.Context) error {
	authHeader := c.Request().Header.Get("Authorization")
	_, err := util.ParseJWTToken(authHeader)
	if err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.Unauthorized, err).Send(c)
	}

	payload := new(pkgdto.SearchGetRequest)
	if err := c.Bind(payload); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err).Send(c)
	}

	if err := c.Validate(payload); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.Validation, err).Send(c)
	}

	result, err := h.service.Find(c.Request().Context(), payload)
	if err != nil {
		return res.ErrorResponse(err).Send(c)
	}

	return res.CustomSuccessBuilder(http.StatusOK, result.Data, "Get users success", &result.PaginationInfo).Send(c)
}

func (h *handler) GetById(c echo.Context) error {
	authHeader := c.Request().Header.Get("Authorization")
	_, err := util.ParseJWTToken(authHeader)
	if err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.Unauthorized, err).Send(c)
	}

	payload := new(pkgdto.ByIDRequest)
	if err := c.Bind(payload); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err).Send(c)
	}
	if err := c.Validate(payload); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.Validation, err).Send(c)
	}

	result, err := h.service.FindByID(c.Request().Context(), payload)
	if err != nil {
		return res.ErrorResponse(err).Send(c)
	}

	return res.SuccessResponse(result).Send(c)
}

func (h *handler) UpdateById(c echo.Context) error {
	payload := new(dto.UpdateUserRequestBody)
	if err := c.Bind(payload); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err).Send(c)
	}
	if err := c.Validate(payload); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.Validation, err).Send(c)
	}
	authHeader := c.Request().Header.Get("Authorization")
	jwtClaims, err := util.ParseJWTToken(authHeader)
	isSameUser := jwtClaims.UserID == *payload.ID
	log.Println(isSameUser)
	if (err != nil) || !isSameUser {
		return res.ErrorBuilder(&res.ErrorConstant.Unauthorized, err).Send(c)
	}
	result, err := h.service.UpdateById(c.Request().Context(), payload)
	if err != nil {
		return res.ErrorResponse(err).Send(c)
	}

	return res.SuccessResponse(result).Send(c)
}

func (h *handler) DeleteById(c echo.Context) error {
	payload := new(pkgdto.ByIDRequest)
	if err := c.Bind(payload); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err).Send(c)
	}
	if err := c.Validate(payload); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.Validation, err).Send(c)
	}
	authHeader := c.Request().Header.Get("Authorization")
	jwtClaims, err := util.ParseJWTToken(authHeader)
	if (err != nil) || jwtClaims.UserID != payload.ID {
		return res.ErrorBuilder(&res.ErrorConstant.Unauthorized, err).Send(c)
	}
	result, err := h.service.DeleteById(c.Request().Context(), payload)
	if err != nil {
		return res.ErrorResponse(err).Send(c)
	}

	return res.SuccessResponse(result).Send(c)
}
