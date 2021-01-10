package worker

import (
	"github.com/roomBoking/pkg/model"
	"time"
)

type BookingWorker struct{

}

func (b BookingWorker) AddBooking(booking model.Booking) (int, error) {
	row := db.QueryRow(`INSERT INTO booking (roomid, startdate, finishdate) VALUES ($1,$2,$3) returning id`,booking.RoomId, booking.StartDateTime, booking.FinishDateTime)
	var id int
	err := row.Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (b BookingWorker) DeleteBooking(bookingId int) error {
	_, err := db.Exec(`DELETE FROM booking WHERE id =$1`, bookingId)
	if err != nil {
		return err
	}
	return  nil
}

func (b BookingWorker) GetBookingList(roomId int) ([]model.ReturnBooking, error) {
	rows, err := db.Query(`SELECT id, startdate, finishdate FROM booking where roomid = $1 ORDER BY startdate`, roomId)
	if err != nil {
		return nil,err
	}
	defer rows.Close()
	bookinglist := make([]model.ReturnBooking, 0)
	for rows.Next() {
		var startTime, finishTime time.Time
		booking := model.ReturnBooking{}
		err := rows.Scan(&booking.Id, &startTime, &finishTime)
		if err != nil {
			return nil, err
		}
		booking.StartDateTime = startTime.Format("2006-01-02")
		booking.FinishDateTime = finishTime.Format("2006-01-02")
		bookinglist = append(bookinglist, booking)
	}
	return  bookinglist, nil
}


