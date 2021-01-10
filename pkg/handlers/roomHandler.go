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

func (h Hotel) AddHotelRoom(w http.ResponseWriter, r *http.Request)  {
	room := model.Room{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Warning("Read request error")
	}
	err = json.Unmarshal(body, &room)
	if err != nil {
		log.Warning("Unmarshal error")
	}
	room.Id, room.Error = h.Rooms.AddRoom(room)
	if room.Error != nil{
		err := returnJson("error", error.Error(room.Error))
		log.Info("Error adding room: ", room)
		w.Write(err)
		return
	}
	log.Info("Adding room: ", room)
	w.Write(returnJson("id", strconv.Itoa(room.Id)))
}

func (h Hotel) DeleteHotelRoom(w http.ResponseWriter, r *http.Request)  {
	paramFromURL := mux.Vars(r)
	id, _ := strconv.Atoi(paramFromURL["id"])
	err := h.Rooms.DeleteRoom(id)
	if err != nil {
		log.Info("error deleting room")
		w.Write(returnJson("error", error.Error(err)))
		return
	}
	log.Info("Room was deleted")
	w.Write(returnJson("status", "ok"))
}

func (h Hotel) GetHotelRoomList(w http.ResponseWriter, r *http.Request)  {
	sortDate, _ := strconv.ParseBool(r.URL.Query().Get("date"))
	sortPrice, _ := strconv.ParseBool(r.URL.Query().Get("price"))
	direction, _ := strconv.ParseBool(r.URL.Query().Get("direction"))
	roomList, err := h.Rooms.GetRoomList(sortDate, sortPrice, direction)
	if err != nil {
		log.Info("error getting room list")
		w.Write(returnJson("error", error.Error(err)))
		return
	}
	jsonData , err := json.Marshal(roomList)
	w.Write(jsonData)
}