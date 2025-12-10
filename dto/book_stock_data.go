package dto

type CreateBookStockRequest struct {
	BookId string   `json:"book_id" validate:"required"`
	Codes  []string `json:"codes" validate:"required,min=1,unique"`
}

type DeleteBookStockRequest struct {
	Codes []string // ambil dari query
}
