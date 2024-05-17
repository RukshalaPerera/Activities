package model

import "time"

type Reservation struct {
	Id          uint      `gorm:"primary_key;unique;not null;auto_increment" json:"id"`
	BookId      uint      `gorm:"not null" json:"book_id"` // Foreign key - book
	UserId      uint      `gorm:"not null" json:"user_id"` // Foreign key - user
	IsCompleted bool      `gorm:"default:false" json:"is_completed"`
	EndTime     time.Time `json:"end_time"`
	StartTime   time.Time `json:"start_time"`
	Status      string    `json:"status"`
}

func NewReservation(bookId uint, userId uint, isCompleted bool, end_time, start_time time.Time, status string) *Reservation {
	return &Reservation{
		BookId:      bookId,
		UserId:      userId,
		IsCompleted: isCompleted,
		EndTime:     end_time,
		StartTime:   start_time,
		Status:      status,
	}
}

func (r *Reservation) Completed() bool {
	return r.IsCompleted
}
