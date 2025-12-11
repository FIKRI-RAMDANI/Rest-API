package domain

import (
	"context"
	"database/sql"

	"github.com/FIKRI-RAMDANI/Rest-API/dto"
)

const (
	JournalStatusInProgress  = "In_Progress"
	JournalStatusInCompleted = "In_Completed"
)

type Journal struct {
	Id         string       `db:"id"`
	BookId     string       `db:"boo_id"`
	StockCode  string       `db:"stock_code"`
	CustomerId string       `db:"customer_id"`
	Status     string       `db:"status"`
	BorrowedAt sql.NullTime `db:"borrowed_at"`
	Returned   sql.NullTime `db:"returned_at"`
}

type JournalSearch struct {
	CustomerId string
	Status     string
}

type JournalRepository interface {
	Find(ctx context.Context, se JournalSearch) ([]Journal, error)
	FindId(ctx context.Context, id string) (Journal, error)
	Save(ctx context.Context, journal *Journal) error
	Update(ctx context.Context, journal *Journal) error
}

type JournalService interface {
	Index(ctx context.Context, se JournalSearch) ([]dto.JournalData, error)
	Create(ctx context.Context, req dto.CreateJournalRequest) error
	Return(ctx context.Context, req dto.ReturnJournalRequest) error
}
