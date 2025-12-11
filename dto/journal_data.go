package dto

import "time"

type JournalData struct {
	Id         string       `json:"id"`
	BookStock  string       `json:"bookStock"`
	Book       BookData     `json:"book"`
	Customer   CustomerData `json:"customer"`
	Status     string       `json:"status"`
	BorrowedAt time.Time    `json:"borrowedAt"`
	ReturnedAt time.Time    `json:"returnedAt"`
}

type CreateJournalRequest struct {
	BookId     string `json:"book_id"`
	BookStock  string `json:"book_stock"`
	CustomerId string `json:"customer_id"`
}

type ReturnJournalRequest struct {
	JournalId string `json:"journal_id"`
}
