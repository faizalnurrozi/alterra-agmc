package auth

import (
	"context"
	"errors"
	"github.com/faizalnurrozi/alterra-agmc/internal/dto"
	"github.com/faizalnurrozi/alterra-agmc/internal/factory"
	"github.com/faizalnurrozi/alterra-agmc/internal/pkg/util"
	"github.com/faizalnurrozi/alterra-agmc/internal/repository"
	"github.com/faizalnurrozi/alterra-agmc/pkg/constant"
	pkgutil "github.com/faizalnurrozi/alterra-agmc/pkg/util"
	res "github.com/faizalnurrozi/alterra-agmc/pkg/util/response"
)

type service struct {
	UserRepository repository.User
}

type Service interface {
	LoginByEmailAndPassword(ctx context.Context, payload *dto.ByEmailAndPasswordRequest) (*dto.UserWithJWTResponse, error)
	RegisterByEmailAndPassword(ctx context.Context, payload *dto.RegisterUserRequestBody) (*dto.UserWithJWTResponse, error)
}

func NewService(f *factory.Factory) Service {
	return &service{
		UserRepository: f.UserRepository,
	}
}

func (s *service) LoginByEmailAndPassword(ctx context.Context, payload *dto.ByEmailAndPasswordRequest) (*dto.UserWithJWTResponse, error) {
	var result *dto.UserWithJWTResponse

	data, err := s.UserRepository.FindByEmail(ctx, &payload.Email)
	if err != nil {
		if err == constant.RecordNotFound {
			return result, res.ErrorBuilder(&res.ErrorConstant.NotFound, err)
		}
		return result, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	if !(pkgutil.CompareHashPassword(payload.Password, data.Password)) {
		return result, res.ErrorBuilder(
			&res.ErrorConstant.EmailOrPasswordIncorrect,
			errors.New(res.ErrorConstant.EmailOrPasswordIncorrect.Response.Meta.Message),
		)
	}

	claims := util.CreateJWTClaims(data.Email, data.ID)
	token, err := util.CreateJWTToken(claims)
	if err != nil {
		return result, res.ErrorBuilder(
			&res.ErrorConstant.InternalServerError,
			errors.New("error when generating token"),
		)
	}

	result = &dto.UserWithJWTResponse{
		UserResponse: dto.UserResponse{
			ID:         data.ID,
			Name:       data.Name,
			Email:      data.Email,
			Gender:     data.Gender,
			Nik:        data.Nik,
			BirthDate:  data.BirthDate.String(),
			YearOfJoin: data.YearOfJoin,
		},
		JWT: token,
	}

	return result, nil
}

func (s *service) RegisterByEmailAndPassword(ctx context.Context, payload *dto.RegisterUserRequestBody) (*dto.UserWithJWTResponse, error) {
	var result *dto.UserWithJWTResponse
	isExist, err := s.UserRepository.ExistByEmail(ctx, &payload.Email)
	if err != nil {
		return result, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}
	if isExist {
		return result, res.ErrorBuilder(&res.ErrorConstant.Duplicate, errors.New("user already exists"))
	}

	hashedPassword, err := pkgutil.HashPassword(payload.Password)
	if err != nil {
		return result, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}
	payload.Password = hashedPassword

	data, err := s.UserRepository.Save(ctx, payload)
	if err != nil {
		return result, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	claims := util.CreateJWTClaims(data.Email, data.ID)
	token, err := util.CreateJWTToken(claims)
	if err != nil {
		return result, res.ErrorBuilder(
			&res.ErrorConstant.InternalServerError,
			errors.New("error when generating token"),
		)
	}

	result = &dto.UserWithJWTResponse{
		UserResponse: dto.UserResponse{
			ID:            data.ID,
			Name:          data.Name,
			Email:         data.Email,
			Gender:        data.Gender,
			Nik:           data.Nik,
			BirthDate:     data.BirthDate.String(),
			MarriedStatus: data.MarriedStatus,
			YearOfJoin:    data.YearOfJoin,
		},
		JWT: token,
	}

	return result, nil
}
