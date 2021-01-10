package handlers

import (
	"encoding/json"
	"github.com/roomBoking/pkg/model"
	log "github.com/sirupsen/logrus"
)

type Hotel struct {
	Rooms interface{
		AddRoom(room model.Room) (int, error)
		DeleteRoom(roomId int) error
		GetRoomList(sortDate, sortPrice, direction bool) ([]model.Room, error)
	}
	Booking interface{
		AddBooking(booking model.Booking) (int, error)
		DeleteBooking(bookingId int) error
		GetBookingList(roomId int) ([]model.ReturnBooking, error)
	}
}


func returnJson(name string, value string) []byte {
	jsonData , err := json.Marshal(name +":"+ value)
	if err != nil{
		log.Warning("json.Marshal error")
	}
	return jsonData
}