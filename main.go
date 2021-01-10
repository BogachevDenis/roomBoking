package main

import (

	"github.com/gorilla/mux"
	"github.com/roomBoking/pkg/handlers"
	"github.com/roomBoking/pkg/worker"
	"github.com/joho/godotenv"

	"net/http"
	"os"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
)


func init()  {
	e := godotenv.Load()
	if e != nil {
		log.Fatal(e)
	}
}

func main() {
	port := os.Getenv("PORT")
	pgUser := os.Getenv("POSTGRES_USER")
	pgPass:= os.Getenv("POSTGRES_PASSWORD")
	pgBase := os.Getenv("POSTGRES_DB")

	hotel := handlers.Hotel{
		Rooms: worker.RoomWorker{},
		Booking: worker.BookingWorker{},
	}
	err := worker.Connect(pgUser,pgPass,pgBase)
	if err != nil {
		log.Fatal(err)
	}

	r := mux.NewRouter()
	r.HandleFunc("/room/add", hotel.AddHotelRoom).Methods("POST")
	r.HandleFunc("/room/{id}", hotel.DeleteHotelRoom).Methods("DELETE")
	r.HandleFunc("/room", hotel.GetHotelRoomList).Methods("GET")

	r.HandleFunc("/booking/add", hotel.AddHotelBooking).Methods("POST")
	r.HandleFunc("/booking/{id}", hotel.DeleteHotelBooking).Methods("DELETE")
	r.HandleFunc("/booking", hotel.GetHotelBookingList).Methods("GET")
	log.Info("Starting Server at " + port)
	http.ListenAndServe(":" + port, r)
}
