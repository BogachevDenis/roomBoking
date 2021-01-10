package handlers

import (
	"encoding/json"
	"github.com/joho/godotenv"
	"github.com/roomBoking/pkg/model"
	"github.com/roomBoking/pkg/worker"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	_ "github.com/lib/pq"
)


func TestAddHotelRoom(t *testing.T)  {
	var room model.Room
	type testpair struct {
		value 	string
		notwant 	int
	}
	var tests = []testpair{
		{ `{"price":15, "description":"okoj"}`, 0},
	}
	godotenv.Load("../../.env")
	pgUser := os.Getenv("POSTGRES_USER")
	pgPass:= os.Getenv("POSTGRES_PASSWORD")
	pgBase := os.Getenv("POSTGRES_DB")
	err := worker.Connect(pgUser,pgPass,pgBase)
	if err != nil {
		log.Print(err)
	}
	for _, pair := range tests {
		r := strings.NewReader(pair.value)
		req, err := http.NewRequest("POST", "/room/add", r)
		if err != nil {
			t.Fatal(err)

		}
		rr := httptest.NewRecorder()
		hotel := Hotel{
			Rooms: worker.RoomWorker{},
			Booking: worker.BookingWorker{},
		}
		handler := http.HandlerFunc(hotel.AddHotelRoom)
		handler.ServeHTTP(rr, req)
		body, _ := ioutil.ReadAll(rr.Body)
		json.Unmarshal(body, &room)
		if room.Id > pair.notwant {
			t.Errorf("handler returned unexpected body: got %v want %v",
				room.Id, pair.notwant)
		}

	}
}


func TestDeleteHotelRoom(t *testing.T)  {
	type status struct {
		status string `json:"status"`
	}
	var st status
	type testpair struct {
		value 	string
		want 	string
	}
	var tests = []testpair{
		{ "22", "ok"},
	}
	godotenv.Load("../../.env")
	pgUser := os.Getenv("POSTGRES_USER")
	pgPass:= os.Getenv("POSTGRES_PASSWORD")
	pgBase := os.Getenv("POSTGRES_DB")
	err := worker.Connect(pgUser,pgPass,pgBase)
	if err != nil {
		log.Print(err)
	}
	for _, pair := range tests {
		r := strings.NewReader(pair.value)
		req, err := http.NewRequest("DELETE", "/room/" + pair.value, r)
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		hotel := Hotel{
			Rooms: worker.RoomWorker{},
			Booking: worker.BookingWorker{},
		}
		handler := http.HandlerFunc(hotel.DeleteHotelRoom)
		handler.ServeHTTP(rr, req)
		body, _ := ioutil.ReadAll(rr.Body)
		json.Unmarshal(body, &st)
		if st.status > pair.want {
			t.Errorf("handler returned unexpected body: got %v want %v",
				st.status, pair.want)
		}
	}
}


func TestGetHotelRoomList(t *testing.T)  {
	var room model.Room
	type testpair struct {
		value 	string
		want 	string
	}
	var tests = []testpair{
		{ "date=false&price=true&direction=false", "nil"},
	}
	godotenv.Load("../../.env")
	pgUser := os.Getenv("POSTGRES_USER")
	pgPass:= os.Getenv("POSTGRES_PASSWORD")
	pgBase := os.Getenv("POSTGRES_DB")
	err := worker.Connect(pgUser,pgPass,pgBase)
	if err != nil {
		log.Print(err)
	}
	for _, pair := range tests {
		r := strings.NewReader(pair.value)
		req, err := http.NewRequest("GET", "/room?" + pair.value, r)
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		hotel := Hotel{
			Rooms: worker.RoomWorker{},
			Booking: worker.BookingWorker{},
		}
		handler := http.HandlerFunc(hotel.DeleteHotelRoom)
		handler.ServeHTTP(rr, req)
		body, _ := ioutil.ReadAll(rr.Body)
		json.Unmarshal(body, &room)
		if room.Error != nil {
			t.Errorf("handler returned unexpected body: got %v want %v",
				error.Error(room.Error), pair.want)
		}
	}
}