package dbrepo

import (
	"bookingApp/internal/models"
	"errors"
	"log"
	"time"
)

func (m *testDBRepo) AllUsers() bool {
	return true
}

// InsertReservation inserts a reservation into database
func (m *testDBRepo) InsertReservation(res models.Reservation) (int, error) {
	// If the room id is 2, then fail; otherwise, pass
	if res.RoomID == 2 {
		return 0, errors.New("insert reservation error")
	}
	return 1, nil
}

// InsertRoomRestriction inserts a room restriction into database
func (m *testDBRepo) InsertRoomRestriction(r models.RoomsRestriction) error {
	if r.RoomID == 1000 {
		return errors.New("insert room restriction error")
	}
	return nil
}

// SearchAvailabilityByDatesByRoomID Returns true if availability exists
func (m *testDBRepo) SearchAvailabilityByDatesByRoomID(start, end time.Time, roomID int) (bool, error) {
	layout := "2006-01-02"
	str := "2049-12-31"
	t, err := time.Parse(layout, str)
	if err != nil {
		log.Println(err)
	}

	testDateToFail, err := time.Parse(layout, "2060-01-01")
	if err != nil {
		log.Println(err)
	}

	if start == testDateToFail {
		return false, errors.New("some error")
	}

	if start.After(t) {
		return false, nil
	}

	return true, nil
}

// SearchAvailabilityForAllRooms returns a slice for available rooms if any
func (m *testDBRepo) SearchAvailabilityForAllRooms(start, end time.Time) ([]models.Room, error) {
	var rooms []models.Room

	layout := "2006-01-02"
	str := "2049-12-31"
	t, err := time.Parse(layout, str)
	if err != nil {
		log.Println(err)
	}

	testDateToFail, err := time.Parse(layout, "2060-01-01")
	if err != nil {
		log.Println(err)
	}

	if start == testDateToFail {
		return rooms, errors.New("some error")
	}

	if start.After(t) {
		return rooms, nil
	}

	room := models.Room{
		ID: 1,
	}
	rooms = append(rooms, room)

	return rooms, nil
}

// GetRoomByID gets a room by ID
func (m *testDBRepo) GetRoomByID(id int) (models.Room, error) {
	var room models.Room
	if id > 2 {
		return room, errors.New("Some error")
	}
	return room, nil
}

func (m *testDBRepo) GetUserByID(id int) (models.User, error) {
	var u models.User
	return u, nil
}

func (m *testDBRepo) UpdateUser(u models.User) error {
	return nil
}

func (m *testDBRepo) Authenticate(email, testPassword string) (int, string, error) {
	if email == "me@here.com" {
		return 1, "", nil
	}
	return 0, "", errors.New("some error")
}

func (m *testDBRepo) AllReservations() ([]models.Reservation, error) {
	var reservations []models.Reservation
	return reservations, nil
}

func (m *testDBRepo) AllNewReservations() ([]models.Reservation, error) {
	var reservations []models.Reservation
	return reservations, nil
}

func (m *testDBRepo) GetReservationByID(id int) (models.Reservation, error) {
	var res models.Reservation
	return res, nil
}

func (m *testDBRepo) UpdateReservation(r models.Reservation) error {
	return nil
}

func (m *testDBRepo) DeleteReservation(id int) error {
	return nil
}

func (m *testDBRepo) UpdateProcessedForReservation(id int, processed int) error {
	return nil
}

func (m *testDBRepo) AllRooms() ([]models.Room, error) {
	var rooms []models.Room
	return rooms, nil
}

func (m *testDBRepo) GetRestrictionForRoomByDate(roomID int, start, end time.Time) ([]models.RoomsRestriction, error) {
	var restrictions []models.RoomsRestriction
	return restrictions, nil
}

func (m *testDBRepo) InsertBlockForRoom(id int, startDate time.Time) error {
	return nil
}

func (m *testDBRepo) DeleteBlockByID(id int) error {
	return nil
}
