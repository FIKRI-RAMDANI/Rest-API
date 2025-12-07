package repository

import (
	"context"
	"database/sql"

	"github.com/FIKRI-RAMDANI/Rest-API/domain"
	"github.com/doug-martin/goqu/v9"
)

type bookStockRepository struct {
	db *goqu.Database
}

func NewStock(con *sql.DB) domain.BookStockRepository {
	return &bookStockRepository{
		db: goqu.New("default", con),
	}
}

func (b bookStockRepository) FindById(ctx context.Context, id string) (stock []domain.BookStock, err error) {
	dataset := b.db.From("book_stocks").Where(goqu.C("id").Eq(id))
	err = dataset.ScanStructsContext(ctx, &stock)
	return
}

func (b bookStockRepository) FindByBookAndCode(ctx context.Context, id string, code string) (stock domain.BookStock, err error) {
	dataset := b.db.From("book_stocks").
		Where(goqu.C("id").Eq(id), goqu.C("code").Eq(code))
	_, err = dataset.ScanStructContext(ctx, &stock)
	return
}

func (b bookStockRepository) Save(ctx context.Context, data []domain.BookStock) error {
	executor := b.db.Insert("book_stocks").Rows(data).Executor()
	_, err := executor.ExecContext(ctx)
	return err
}

func (b bookStockRepository) Update(ctx context.Context, stock *domain.BookStock) error {
	executor := b.db.Update("book_stocks").Where(goqu.C("id").Eq(stock.Code)).Set(stock).Executor()
	_, err := executor.ExecContext(ctx)
	return err
}

func (b bookStockRepository) DeleteByBookID(ctx context.Context, id string) error {
	executor := b.db.Delete("book_stocks").Where(goqu.C("id").Eq(id)).Executor()
	_, err := executor.ExecContext(ctx)
	return err
}

func (b bookStockRepository) DeleteByCode(ctx context.Context, codes []string) error {
	executor := b.db.Delete("book_stocks").Where(goqu.C("code").In(codes)).Executor()
	_, err := executor.ExecContext(ctx)
	return err
}
