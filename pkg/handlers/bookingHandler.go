package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/roomBoking/pkg/model"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"strconv"
)

func (h Hotel) AddHotelBooking(w http.ResponseWriter, r *http.Request)  {
	booking := model.Booking{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Warning("Read request error")
	}
	err = json.Unmarshal(body, &booking)
	if err != nil {
		log.Warning("Unmarshal error")
	}
	booking.ConvertDate()
	if booking.Error != nil{
		err := returnJson("error", error.Error(booking.Error))
		log.Info("Date is not valid")
		w.Write(err)
		return
	}
	booking.Id, booking.Error = h.Booking.AddBooking(booking)
	if booking.Error != nil{
		err := returnJson("error", error.Error(booking.Error))
		log.Info("add booking error")
		w.Write(err)
		return
	}
	w.Write(returnJson("id", strconv.Itoa(booking.Id)))
}

func (h Hotel) DeleteHotelBooking(w http.ResponseWriter, r *http.Request)  {
	paramFromURL := mux.Vars(r)
	id, _ := strconv.Atoi(paramFromURL["id"])
	err := h.Booking.DeleteBooking(id)
	if err != nil {
		log.Info(err)
		deleteError := returnJson("error", error.Error(err))
		w.Write(deleteError)
		return
	}
	w.Write(returnJson("status", "ok"))
}

func (h Hotel) GetHotelBookingList(w http.ResponseWriter, r *http.Request)  {
	roomId, _ := strconv.Atoi(r.URL.Query().Get("roomid"))
	bookingList, err := h.Booking.GetBookingList(roomId)
	if err != nil {
		log.Info("error getting boking list")
		w.Write(returnJson("error", error.Error(err)))
		return
	}
	log.Info("getting boking list", bookingList)
	jsonData , err := json.Marshal(bookingList)
	w.Write(jsonData)
}