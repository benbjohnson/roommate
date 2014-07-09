package roommate_test

import (
	"io/ioutil"
	"log"
	"os"
	"testing"
	"time"

	. "github.com/benbjohnson/roommate"
)

// Ensure that a user can be saved and retrieved.
func TestTx_User(t *testing.T) {
	db := NewTestDB()
	defer db.Close()
	db.Update(func(tx *Tx) error {
		ok(t, tx.SaveUser(&User{ID: "123", Email: "john@gmail.com"}))
		return nil
	})
	db.View(func(tx *Tx) error {
		equals(t, &User{ID: "123", Email: "john@gmail.com"}, tx.User("123"))
		return nil
	})
}

// Ensure that a room can be saved and retrieved.
func TestTx_Room(t *testing.T) {
	db := NewTestDB()
	defer db.Close()
	db.Update(func(tx *Tx) error {
		ok(t, tx.SaveRoom(&Room{ID: 1, Name: "my room"}))
		return nil
	})
	db.View(func(tx *Tx) error {
		equals(t, &Room{ID: 1, Name: "my room"}, tx.Room(1))
		return nil
	})
}

// Ensure that all rooms can be retrieved.
func TestTx_Rooms(t *testing.T) {
	db := NewTestDB()
	defer db.Close()
	db.Update(func(tx *Tx) error {
		ok(t, tx.SaveRoom(&Room{ID: 1, Name: "room 1"}))
		ok(t, tx.SaveRoom(&Room{ID: 2, Name: "room 2"}))
		ok(t, tx.SaveRoom(&Room{ID: 3, Name: "room 3"}))
		return nil
	})
	db.View(func(tx *Tx) error {
		rooms := tx.Rooms()
		equals(t, 3, len(rooms))
		equals(t, "room 1", rooms[0].Name)
		equals(t, "room 2", rooms[1].Name)
		equals(t, "room 3", rooms[2].Name)
		return nil
	})
}

// Ensure that a reservation can be saved and retrieved.
func TestTx_Reservation(t *testing.T) {
	db := NewTestDB()
	defer db.Close()
	db.Update(func(tx *Tx) error {
		ok(t, tx.SaveReservation(&Reservation{ID: 2, RoomID: 1, Time: parsetime("2000-01-01T00:00:00Z"), Duration: time.Second}))
		return nil
	})
	db.View(func(tx *Tx) error {
		equals(t, &Reservation{ID: 2, RoomID: 1, Time: parsetime("2000-01-01T00:00:00Z"), Duration: time.Second}, tx.Reservation(2))
		return nil
	})
}

// TestDB wraps the DB to provide cleanup after closing.
type TestDB struct {
	*DB
}

// NewTestDB creates and opens a new database at a temporary directory.
func NewTestDB() *TestDB {
	db := &TestDB{&DB{}}
	if err := db.Open(tempfile(), 0600); err != nil {
		log.Fatalln("open:", err)
	}
	return db
}

// Close closes and deletes the database.
func (db *TestDB) Close() {
	defer os.RemoveAll(db.Path())
	db.DB.Close()
}

// tempfile returns the path to a non-existent temporary file.
func tempfile() string {
	f, _ := ioutil.TempFile("", "roommate-")
	path := f.Name()
	f.Close()
	os.Remove(path)
	return path
}
