package service

import (
	"context"

	"github.com/FIKRI-RAMDANI/Rest-API/domain"
	"github.com/FIKRI-RAMDANI/Rest-API/dto"
)

type bookStockService struct {
	bookRepository      domain.BookRepository
	bookStockRepository domain.BookStockRepository
}

func NewBookStock(bookRepository domain.BookRepository, bookStockRepository domain.BookStockRepository) domain.BookStockService {
	return &bookStockService{
		bookRepository:      bookRepository,
		bookStockRepository: bookStockRepository,
	}
}

func (b bookStockService) Create(ctx context.Context, req dto.CreateBookStockRequest) error {
	// Cek buku
	book, err := b.bookRepository.FindById(ctx, req.BookId)
	if err != nil {
		return err
	}

	if book.Id == "" {
		return domain.BookNotFound
	}

	stock := make([]domain.BookStock, 0)
	for _, v := range req.Codes {
		stock = append(stock, domain.BookStock{
			Code:   v,
			BookId: req.BookId,
			Status: domain.BookStockStatusAvailable,
		})
	}
	return b.bookStockRepository.Save(ctx, stock)
}

func (b bookStockService) Delete(ctx context.Context, req dto.DeleteBookStockRequest) error {
	return b.bookStockRepository.DeleteByCode(ctx, req.Codes)
}
