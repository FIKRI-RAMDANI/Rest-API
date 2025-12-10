package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/FIKRI-RAMDANI/Rest-API/domain"
	"github.com/doug-martin/goqu/v9"
)

type bookRepository struct {
	db *goqu.Database
}

func NewBook(con *sql.DB) domain.BookRepository {
	return &bookRepository{
		db: goqu.New("default", con),
	}
}

func (br bookRepository) FindAll(ctx context.Context) (books []domain.Book, err error) {
	dataset := br.db.From("books").Where(goqu.C("deleted_at").IsNull())
	err = dataset.ScanStructsContext(ctx, &books)
	return
}

func (br bookRepository) FindById(ctx context.Context, id string) (book domain.Book, err error) {
	dataset := br.db.From("books").Where(goqu.C("deleted_at").IsNull(), goqu.C("id").Eq(id))
	_, err = dataset.ScanStructContext(ctx, &book)
	return
}

func (br bookRepository) Save(ctx context.Context, book *domain.Book) error {
	executor := br.db.Insert("books").Rows(book).Executor()
	_, err := executor.ExecContext(ctx)
	return err
}

func (br bookRepository) Updated(ctx context.Context, book *domain.Book) error {
	executor := br.db.Update("books").
		Where(goqu.C("id").Eq(book.Id)).
		Set(goqu.Record{
			"isbn":        book.Isbn,
			"title":       book.Title,
			"description": book.Description,
			"updated_at":  time.Now(),
		}).Executor()

	_, err := executor.ExecContext(ctx)
	return err
}

func (br bookRepository) Deleted(ctx context.Context, id string) error {
	executor := br.db.Update("books").
		Where(goqu.C("id").Eq(id)).
		Set(goqu.Record{"deleted_at": sql.NullTime{Valid: true, Time: time.Now()}}).Executor()

	_, err := executor.ExecContext(ctx)
	return err
}
