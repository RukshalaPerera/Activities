package model

type Book struct {
	Title           string `gorm:"type:varchar(255)" json:"title"`
	Author          string `gorm:"type:varchar(255)" json:"author"`
	Id              uint64 `gorm:"primaryKey" json:"id"`
	BookName        string `gorm:"type:varchar(255)" json:"book_name"`
	IsBookAvailable bool   `gorm:"type:bool" json:"is_book_available"`
}

func NewBook(title string, author string, id uint64, Book_name string, is_book_available bool) *Book {
	return &Book{
		Title:           title,
		Author:          author,
		Id:              id,
		BookName:        Book_name,
		IsBookAvailable: is_book_available,
	}
}
