package worker

import (
	"github.com/roomBoking/pkg/model"
)

type RoomWorker struct{

}

func (r RoomWorker) AddRoom(room model.Room) (int, error) {
	row := db.QueryRow(`INSERT INTO hotelroom (description, price, dateadding) VALUES ($1,$2,now()) returning id`,room.Description, room.Price)
	var id int
	err := row.Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r RoomWorker) DeleteRoom(roomId int) error {
	_, err := db.Exec(`DELETE FROM hotelroom WHERE id =$1`, roomId)
	if err != nil {
		return err
	}
	_, err = db.Exec(`DELETE FROM booking WHERE roomid =$1`, roomId)
	if err != nil {
		return err
	}
	return  nil
}

func (r RoomWorker) GetRoomList(sortDate, sortPrice, direction bool) ([]model.Room,error) {
	query := `SELECT * FROM hotelroom`
	if sortDate {
		if direction {
			query = `SELECT * FROM hotelroom ORDER BY dateadding`
		} else {
			query = `SELECT * FROM hotelroom ORDER BY dateadding DESC`
		}
	}
	if sortPrice {
		if direction {
			query = `SELECT * FROM hotelroom ORDER BY price`
		} else {
			query = `SELECT * FROM hotelroom ORDER BY price DESC`
		}
	}
	rows, err := db.Query(query)
	if err != nil {
		return nil,err
	}
	defer rows.Close()
	rooms := make([]model.Room, 0)
	for rows.Next() {
		room := model.Room{}
		err := rows.Scan(&room.Id,&room.Description,&room.Price, &room.AddingDate)
		if err != nil {
			return nil, err
		}
		rooms = append(rooms, room)
	}
	return  rooms, nil
}