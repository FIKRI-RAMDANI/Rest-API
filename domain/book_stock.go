package domain

import (
	"context"
	"database/sql"

	"github.com/FIKRI-RAMDANI/Rest-API/dto"
)

const (
	BookStockStatusAvailable = "AVAILABLE"
	BookStockStatusBorrowed  = "BORROWED"
)

type BookStock struct {
	Code       string         `db:"code"`
	BookId     string         `db:"book_id"`
	Status     string         `db:"status"`
	BorrowerId sql.NullString `db:"borrower_id"`
	BorrowerAt sql.NullTime   `db:"borrowed_at"`
}

type BookStockRepository interface {
	FindById(ctx context.Context, id string) ([]BookStock, error)
	FindByBookAndCode(ctx context.Context, id string, code string) (BookStock, error)
	Save(ctx context.Context, data []BookStock) error
	Update(ctx context.Context, stock *BookStock) error
	DeleteByBookID(ctx context.Context, id string) error
	DeleteByCode(ctx context.Context, codes []string) error
}

type BookStockService interface {
	Create(ctx context.Context, req dto.CreateBookStockRequest) error
	Delete(ctx context.Context, req dto.DeleteBookStockRequest) error
}
