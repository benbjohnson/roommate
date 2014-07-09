package roommate

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"os"
	"time"

	"github.com/boltdb/bolt"
)

// DB represents the data layer of the application.
type DB struct {
	*bolt.DB
}

// Open opens and initializes the database.
func (db *DB) Open(path string, mode os.FileMode) error {
	d, err := bolt.Open(path, mode, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return err
	}
	db.DB = d

	// Initialize all the required buckets.
	return db.Update(func(tx *Tx) error {
		tx.CreateBucketIfNotExists([]byte("Users"))
		tx.CreateBucketIfNotExists([]byte("Rooms"))
		tx.CreateBucketIfNotExists([]byte("Reservations"))
		return nil
	})
}

// View executes a function in the context of a read-only transaction.
func (db *DB) View(fn func(*Tx) error) error {
	return db.DB.View(func(tx *bolt.Tx) error {
		return fn(&Tx{tx})
	})
}

// Update executes a function in the context of a writable transaction.
func (db *DB) Update(fn func(*Tx) error) error {
	return db.DB.Update(func(tx *bolt.Tx) error {
		return fn(&Tx{tx})
	})
}

// Tx represents a transaction.
type Tx struct {
	*bolt.Tx
}

// User retrieves a user by id.
func (tx *Tx) User(id string) *User {
	v := tx.Bucket([]byte("Users")).Get([]byte(id))
	if v == nil {
		return nil
	}
	u := &User{}
	_ = json.Unmarshal(v, u)
	return u
}

// SaveUser persist a user to the database.
func (tx *Tx) SaveUser(u *User) error {
	if u == nil {
		return errors.New("user required")
	} else if u.ID == "" {
		return errors.New("user id required")
	}
	b, _ := json.Marshal(u)
	return tx.Bucket([]byte("Users")).Put([]byte(u.ID), b)
}

// Room retrieves a room by id.
func (tx *Tx) Room(id int) *Room {
	v := tx.Bucket([]byte("Rooms")).Get(itob(id))
	if v == nil {
		return nil
	}
	r := &Room{}
	_ = json.Unmarshal(v, r)
	return r
}

// Rooms retrieves all the rooms in the database.
func (tx *Tx) Rooms() []*Room {
	var a []*Room
	tx.Bucket([]byte("Rooms")).ForEach(func(_, v []byte) error {
		r := &Room{}
		_ = json.Unmarshal(v, r)
		a = append(a, r)
		return nil
	})
	return a
}

// SaveRoom persist a room to the database.
func (tx *Tx) SaveRoom(r *Room) error {
	if r == nil {
		return errors.New("room required")
	} else if r.ID == 0 {
		return errors.New("room id required")
	}
	b, _ := json.Marshal(r)
	return tx.Bucket([]byte("Rooms")).Put(itob(r.ID), b)
}

// DeleteRoom removes a room from the database.
func (tx *Tx) DeleteRoom(id int) error {
	// TODO(benbjohnson): Delete all reservations too.
	return tx.Bucket([]byte("Rooms")).Delete(itob(id))
}

// Reservation retrieves a reservation by id.
func (tx *Tx) Reservation(id int) *Reservation {
	v := tx.Bucket([]byte("Reservations")).Get(itob(id))
	if v == nil {
		return nil
	}
	r := &Reservation{}
	_ = json.Unmarshal(v, r)
	return r
}

// SaveReservation persist a reservation to the database.
func (tx *Tx) SaveReservation(r *Reservation) error {
	if r == nil {
		return errors.New("reservation required")
	} else if r.ID == 0 {
		return errors.New("reservation id required")
	}
	b, _ := json.Marshal(r)
	return tx.Bucket([]byte("Reservations")).Put(itob(r.ID), b)
}

// itob converts an int into an 8-byte byte slice.
func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}

// btoi converts an 8-byte byte slice into an int.
func btoi64(b []byte) int {
	return int(binary.BigEndian.Uint64(b))
}
