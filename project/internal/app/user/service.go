package user

import (
	"context"
	"github.com/faizalnurrozi/alterra-agmc/internal/dto"
	"github.com/faizalnurrozi/alterra-agmc/internal/factory"
	"github.com/faizalnurrozi/alterra-agmc/internal/repository"
	"github.com/faizalnurrozi/alterra-agmc/pkg/constant"
	pkgdto "github.com/faizalnurrozi/alterra-agmc/pkg/dto"
	res "github.com/faizalnurrozi/alterra-agmc/pkg/util/response"
)

type service struct {
	UserRepository repository.User
}

type Service interface {
	Find(ctx context.Context, payload *pkgdto.SearchGetRequest) (*pkgdto.SearchGetResponse[dto.UserResponse], error)
	FindByID(ctx context.Context, payload *pkgdto.ByIDRequest) (*dto.UserDetailResponse, error)
	UpdateById(ctx context.Context, payload *dto.UpdateUserRequestBody) (*dto.UserDetailResponse, error)
	DeleteById(ctx context.Context, payload *pkgdto.ByIDRequest) (*dto.UserWithCUDResponse, error)
}

func NewService(f *factory.Factory) Service {
	return &service{
		UserRepository: f.UserRepository,
	}
}

func (s *service) Find(ctx context.Context, payload *pkgdto.SearchGetRequest) (*pkgdto.SearchGetResponse[dto.UserResponse], error) {
	users, info, err := s.UserRepository.FindAll(ctx, payload, &payload.Pagination)
	if err != nil {
		return nil, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	var data []dto.UserResponse

	for _, user := range users {
		data = append(data, dto.UserResponse{
			ID:            user.ID,
			Name:          user.Name,
			Email:         user.Email,
			Gender:        user.Gender,
			Nik:           user.Nik,
			BirthDate:     user.BirthDate.String(),
			MarriedStatus: user.MarriedStatus,
			YearOfJoin:    user.YearOfJoin,
		})

	}

	result := new(pkgdto.SearchGetResponse[dto.UserResponse])
	result.Data = data
	result.PaginationInfo = *info

	return result, nil
}

func (s *service) FindByID(ctx context.Context, payload *pkgdto.ByIDRequest) (*dto.UserDetailResponse, error) {
	data, err := s.UserRepository.FindByID(ctx, payload.ID)
	if err != nil {
		if err == constant.RecordNotFound {
			return &dto.UserDetailResponse{}, res.ErrorBuilder(&res.ErrorConstant.NotFound, err)
		}
		return &dto.UserDetailResponse{}, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	result := &dto.UserDetailResponse{
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
	}

	return result, nil
}

func (s *service) UpdateById(ctx context.Context, payload *dto.UpdateUserRequestBody) (*dto.UserDetailResponse, error) {
	user, err := s.UserRepository.FindByID(ctx, *payload.ID)
	if err != nil {
		if err == constant.RecordNotFound {
			return &dto.UserDetailResponse{}, res.ErrorBuilder(&res.ErrorConstant.NotFound, err)
		}
		return &dto.UserDetailResponse{}, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	_, err = s.UserRepository.Edit(ctx, &user, payload)
	if err != nil {
		return &dto.UserDetailResponse{}, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	result := &dto.UserDetailResponse{
		UserResponse: dto.UserResponse{
			ID:            user.ID,
			Name:          user.Name,
			Email:         user.Email,
			Gender:        user.Gender,
			Nik:           user.Nik,
			BirthDate:     user.BirthDate.String(),
			MarriedStatus: user.MarriedStatus,
			YearOfJoin:    user.YearOfJoin,
		},
	}

	return result, nil
}

func (s *service) DeleteById(ctx context.Context, payload *pkgdto.ByIDRequest) (*dto.UserWithCUDResponse, error) {
	user, err := s.UserRepository.FindByID(ctx, payload.ID)
	if err != nil {
		if err == constant.RecordNotFound {
			return &dto.UserWithCUDResponse{}, res.ErrorBuilder(&res.ErrorConstant.NotFound, err)
		}
		return &dto.UserWithCUDResponse{}, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}
	_, err = s.UserRepository.Destroy(ctx, &user)
	if err != nil {
		return &dto.UserWithCUDResponse{}, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	result := &dto.UserWithCUDResponse{
		UserResponse: dto.UserResponse{
			ID:            user.ID,
			Name:          user.Name,
			Email:         user.Email,
			Gender:        user.Gender,
			Nik:           user.Nik,
			BirthDate:     user.BirthDate.String(),
			MarriedStatus: user.MarriedStatus,
			YearOfJoin:    user.YearOfJoin,
		},
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		DeletedAt: user.DeletedAt,
	}

	return result, nil
}
