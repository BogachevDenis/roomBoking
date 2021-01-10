package model

import (
	"errors"
	"time"
)

type Room struct{
	Id 			int		`json:"id"`
	Description string	`json:"description"`
	Price 		int		`json:"price"`
	AddingDate	time.Time
	Error 		error	`json:"error"`
}

type Booking struct{
	Id 				int			`json:"id"`
	RoomId 			int			`json:"roomid"`
	StartDate 		string		`json:"startdate"`
	FinishDate		string		`json:"finishdate"`
	StartDateTime 	time.Time
	FinishDateTime	time.Time
	Error 			error		`json:"error"`
}

type ReturnBooking struct {
	Id				int 	`json:"id"`
	StartDateTime 	string	`json:"startdate"`
	FinishDateTime	string	`json:"finishdate"`
}

func (b *Booking) ConvertDate() {
	t := time.Now()
	b.StartDateTime, _ = time.Parse("2006-01-02", b.StartDate)
	b.FinishDateTime, _ = time.Parse("2006-01-02", b.FinishDate)
	if b.StartDateTime.Before(t) || b.FinishDateTime.Before(t) || b.FinishDateTime.Before(b.StartDateTime)  {
		b.Error = errors.New("time is not valid")
	}
}