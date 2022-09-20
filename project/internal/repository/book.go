package repository

import (
	"context"
	"github.com/faizalnurrozi/alterra-agmc/internal/dto"
	"github.com/faizalnurrozi/alterra-agmc/internal/model"
	pkgdto "github.com/faizalnurrozi/alterra-agmc/pkg/dto"
	"gorm.io/gorm"
	"strings"
	"time"
)

type Book interface {
	FindAll(ctx context.Context, payload *pkgdto.SearchGetRequest, pagination *pkgdto.Pagination) ([]model.Book, *pkgdto.PaginationInfo, error)
	FindByID(ctx context.Context, id uint) (model.Book, error)
	Save(ctx context.Context, book *dto.CreateBookRequestBody) (model.Book, error)
	Edit(ctx context.Context, oldBook *model.Book, updateData *dto.UpdateBookRequestBody) (*model.Book, error)
	Destroy(ctx context.Context, book *model.Book) (*model.Book, error)
	ExistByTitle(ctx context.Context, title string) (bool, error)
}

type book struct {
	Db *gorm.DB
}

func NewBookRepository(db *gorm.DB) *book {
	return &book{
		db,
	}
}

func (r *book) FindAll(ctx context.Context, payload *pkgdto.SearchGetRequest, pagination *pkgdto.Pagination) ([]model.Book, *pkgdto.PaginationInfo, error) {
	var book []model.Book
	var count int64

	query := r.Db.WithContext(ctx).Model(&model.Book{})

	if payload.Search != "" {
		search := "%" + strings.ToLower(payload.Search) + "%"
		query = query.Where("lower(title) LIKE ?", search)
	}

	countQuery := query
	if err := countQuery.Count(&count).Error; err != nil {
		return nil, nil, err
	}

	limit, offset := pkgdto.GetLimitOffset(pagination)

	err := query.Limit(limit).Offset(offset).Find(&book).Error

	return book, pkgdto.CheckInfoPagination(pagination, count), err
}

func (r *book) FindByID(ctx context.Context, id uint) (model.Book, error) {
	var book model.Book
	if err := r.Db.WithContext(ctx).Model(&model.Book{}).Where("id = ?", id).First(&book).Error; err != nil {
		return book, err
	}
	return book, nil
}

func (r *book) Save(ctx context.Context, user *dto.CreateBookRequestBody) (model.Book, error) {
	dateString := *user.DatePublished
	DatePublished, _ := time.Parse("2006-01-02", dateString)
	newUser := model.Book{
		Title:         *user.Title,
		Isbn:          *user.Isbn,
		Author:        *user.Author,
		Publisher:     *user.Publisher,
		DatePublished: DatePublished,
		StatusDisplay: *user.StatusDisplay,
	}

	if err := r.Db.WithContext(ctx).Save(&newUser).Error; err != nil {
		return newUser, err
	}
	return newUser, nil
}

func (r *book) Edit(ctx context.Context, oldBook *model.Book, updateData *dto.UpdateBookRequestBody) (*model.Book, error) {
	if updateData.Title != nil {
		oldBook.Title = *updateData.Title
	}
	if updateData.Isbn != nil {
		oldBook.Isbn = *updateData.Isbn
	}
	if updateData.Author != nil {
		oldBook.Author = *updateData.Author
	}
	if updateData.Publisher != nil {
		oldBook.Publisher = *updateData.Publisher
	}
	if updateData.DatePublished != nil {
		dateString := *updateData.DatePublished
		DatePublished, _ := time.Parse("2006-01-02", dateString)
		oldBook.DatePublished = DatePublished
	}
	if updateData.StatusDisplay != nil {
		oldBook.StatusDisplay = *updateData.StatusDisplay
	}

	if err := r.Db.WithContext(ctx).Save(oldBook).Find(oldBook).Error; err != nil {
		return nil, err
	}

	return oldBook, nil
}

func (r *book) Destroy(ctx context.Context, book *model.Book) (*model.Book, error) {
	if err := r.Db.WithContext(ctx).Delete(book).Error; err != nil {
		return nil, err
	}
	return book, nil
}

func (r *book) ExistByTitle(ctx context.Context, title string) (bool, error) {
	var (
		count   int64
		isExist bool
	)
	if err := r.Db.WithContext(ctx).Model(&model.Book{}).Where("title = ?", title).Count(&count).Error; err != nil {
		return isExist, err
	}
	if count > 0 {
		isExist = true
	}
	return isExist, nil
}
