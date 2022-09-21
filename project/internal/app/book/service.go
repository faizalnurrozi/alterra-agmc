package book

import (
	"context"
	"github.com/faizalnurrozi/alterra-agmc/internal/dto"
	"github.com/faizalnurrozi/alterra-agmc/internal/factory"
	"github.com/faizalnurrozi/alterra-agmc/internal/repository"
	"github.com/faizalnurrozi/alterra-agmc/pkg/constant"
	pkgdto "github.com/faizalnurrozi/alterra-agmc/pkg/dto"
	res "github.com/faizalnurrozi/alterra-agmc/pkg/util/response"
	"github.com/pkg/errors"
)

type service struct {
	BookRepository repository.Book
}

type Service interface {
	Find(ctx context.Context, payload *pkgdto.SearchGetRequest) (*pkgdto.SearchGetResponse[dto.BookResponse], error)
	FindByID(ctx context.Context, payload *pkgdto.ByIDRequest) (*dto.BookResponse, error)
	Store(ctx context.Context, payload *dto.CreateBookRequestBody) (*dto.BookResponse, error)
	UpdateById(ctx context.Context, payload *dto.UpdateBookRequestBody) (*dto.BookResponse, error)
	DeleteById(ctx context.Context, payload *pkgdto.ByIDRequest) (*dto.BookWithCUDResponse, error)
}

func NewService(f *factory.Factory) Service {
	return &service{
		BookRepository: f.BookRepository,
	}
}

func (s *service) Find(ctx context.Context, payload *pkgdto.SearchGetRequest) (*pkgdto.SearchGetResponse[dto.BookResponse], error) {
	books, info, err := s.BookRepository.FindAll(ctx, payload, &payload.Pagination)
	if err != nil {
		return nil, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	var data []dto.BookResponse

	for _, book := range books {
		data = append(data, dto.BookResponse{
			ID:            book.ID,
			Title:         book.Title,
			Isbn:          book.Isbn,
			Author:        book.Author,
			Publisher:     book.Publisher,
			DatePublished: book.DatePublished,
			StatusDisplay: book.StatusDisplay,
		})

	}

	result := new(pkgdto.SearchGetResponse[dto.BookResponse])
	result.Data = data
	result.PaginationInfo = *info

	return result, nil
}
func (s *service) FindByID(ctx context.Context, payload *pkgdto.ByIDRequest) (*dto.BookResponse, error) {
	var result dto.BookResponse
	data, err := s.BookRepository.FindByID(ctx, payload.ID)
	if err != nil {
		if err == constant.RecordNotFound {
			return &dto.BookResponse{}, res.ErrorBuilder(&res.ErrorConstant.NotFound, err)
		}
		return &dto.BookResponse{}, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	result.ID = data.ID
	result.Title = data.Title
	result.Isbn = data.Isbn
	result.Author = data.Author
	result.Publisher = data.Publisher
	result.DatePublished = data.DatePublished
	result.StatusDisplay = data.StatusDisplay

	return &result, nil
}

func (s *service) Store(ctx context.Context, payload *dto.CreateBookRequestBody) (*dto.BookResponse, error) {
	var result dto.BookResponse
	isExist, err := s.BookRepository.ExistByTitle(ctx, *payload.Title)
	if err != nil {
		return &result, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}
	if isExist {
		return &result, res.ErrorBuilder(&res.ErrorConstant.Duplicate, errors.New("book already exists"))
	}

	data, err := s.BookRepository.Save(ctx, payload)
	if err != nil {
		return &result, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	result.ID = data.ID
	result.Title = data.Title
	result.Isbn = data.Isbn
	result.Author = data.Author
	result.Publisher = data.Publisher
	result.DatePublished = data.DatePublished
	result.StatusDisplay = data.StatusDisplay

	return &result, nil
}

func (s *service) UpdateById(ctx context.Context, payload *dto.UpdateBookRequestBody) (*dto.BookResponse, error) {
	book, err := s.BookRepository.FindByID(ctx, *payload.ID)
	if err != nil {
		if err == constant.RecordNotFound {
			return &dto.BookResponse{}, res.ErrorBuilder(&res.ErrorConstant.NotFound, err)
		}
		return &dto.BookResponse{}, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	_, err = s.BookRepository.Edit(ctx, &book, payload)
	if err != nil {
		return &dto.BookResponse{}, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}
	var result dto.BookResponse
	result.ID = book.ID
	result.Isbn = book.Isbn
	result.Author = book.Author
	result.Publisher = book.Publisher
	result.DatePublished = book.DatePublished
	result.StatusDisplay = book.StatusDisplay

	return &result, nil
}
func (s *service) DeleteById(ctx context.Context, payload *pkgdto.ByIDRequest) (*dto.BookWithCUDResponse, error) {
	book, err := s.BookRepository.FindByID(ctx, payload.ID)
	if err != nil {
		if err == constant.RecordNotFound {
			return &dto.BookWithCUDResponse{}, res.ErrorBuilder(&res.ErrorConstant.NotFound, err)
		}
		return &dto.BookWithCUDResponse{}, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}
	_, err = s.BookRepository.Destroy(ctx, &book)
	if err != nil {
		return &dto.BookWithCUDResponse{}, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	result := &dto.BookWithCUDResponse{
		BookResponse: dto.BookResponse{
			ID:            book.ID,
			Title:         book.Title,
			Isbn:          book.Isbn,
			Author:        book.Author,
			Publisher:     book.Publisher,
			DatePublished: book.DatePublished,
			StatusDisplay: book.StatusDisplay,
		},
		CreatedAt: book.CreatedAt,
		UpdatedAt: book.UpdatedAt,
		DeletedAt: book.DeletedAt,
	}

	return result, nil
}
