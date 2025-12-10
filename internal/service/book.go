package service

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/FIKRI-RAMDANI/Rest-API/domain"
	"github.com/FIKRI-RAMDANI/Rest-API/dto"
	"github.com/google/uuid"
)

type bookService struct {
	bookRepository      domain.BookRepository
	bookstockRepository domain.BookStockRepository
}

func NewBook(bookRepository domain.BookRepository, bookStockRepository domain.BookStockRepository) domain.BookService {
	return &bookService{
		bookRepository:      bookRepository,
		bookstockRepository: bookStockRepository,
	}
}

func (b bookService) Index(ctx context.Context) ([]dto.BookData, error) {
	result, err := b.bookRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	var data []dto.BookData
	for _, v := range result {
		data = append(data, dto.BookData{
			Id:          v.Id,
			Isbn:        v.Isbn,
			Title:       v.Title,
			Description: v.Description,
		})
	}
	return data, nil
}

func (b bookService) Show(ctx context.Context, id string) (dto.BookShowData, error) {
	persisted, err := b.bookRepository.FindById(ctx, id)
	if err != nil {

		return dto.BookShowData{}, err
	}
	if persisted.Id == "" {
		return dto.BookShowData{}, domain.BookNotFound
	}
	stocks, err := b.bookstockRepository.FindById(ctx, persisted.Id)
	if err != nil {
		return dto.BookShowData{}, err
	}
	stocksData := make([]dto.BookStockData, 0)
	for _, v := range stocks {
		stocksData = append(stocksData, dto.BookStockData{
			Code:   v.Code,
			Status: v.Status,
		})
	}
	return dto.BookShowData{
		BookData: dto.BookData{
			Id:          persisted.Id,
			Isbn:        persisted.Isbn,
			Title:       persisted.Title,
			Description: persisted.Description,
		},
		Stocks: stocksData,
	}, nil
}

func (b bookService) Create(ctx context.Context, req dto.CreateBookRequest) error {
	book := domain.Book{
		Id:          uuid.NewString(),
		Isbn:        req.Isbn,
		Title:       req.Title,
		Description: req.Description,
		CreatedAt:   sql.NullTime{Valid: true, Time: time.Now()},
	}
	return b.bookRepository.Save(ctx, &book)
}

func (b bookService) Update(ctx context.Context, req dto.UpdateBookRequest) error {
	persisted, err := b.bookRepository.FindById(ctx, req.Id)
	if err != nil {
		return err
	}
	if persisted.Id == "" {
		return errors.New("book not found")
	}
	persisted.Isbn = req.Isbn
	persisted.Title = req.Title
	persisted.Description = req.Description
	persisted.UpdatedAt = sql.NullTime{Valid: true, Time: time.Now()}

	return b.bookRepository.Updated(ctx, &persisted)
}

func (b bookService) Delete(ctx context.Context, id string) error {
	exist, err := b.bookRepository.FindById(ctx, id)
	if err != nil {
		return err
	}
	if exist.Id == "" {
		return errors.New("book not found")
	}
	err = b.bookRepository.Deleted(ctx, id)
	if err != nil {
		return err
	}
	return b.bookstockRepository.DeleteByBookID(ctx, exist.Id)
}
