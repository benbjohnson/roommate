package roommate

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// NewHandler returns a root HTTP handler for the application.
func NewHandler(db *DB) http.Handler {
	h := &Handler{DB: db}

	r := mux.NewRouter()
	r.HandleFunc("/", h.Root)

	r.HandleFunc("/rooms", h.Rooms).Methods("GET")
	r.HandleFunc("/rooms", h.CreateRoom).Methods("POST")
	r.HandleFunc("/rooms/new", h.NewRoom).Methods("GET")
	r.HandleFunc("/rooms/{id}", h.Room).Methods("GET")
	r.HandleFunc("/rooms/{id}/edit", h.EditRoom).Methods("GET")
	r.HandleFunc("/rooms/{id}", h.UpdateRoom).Methods("POST")
	r.HandleFunc("/rooms/{id}/delete", h.DeleteRoom).Methods("GET")
	return r
}

// Handler represents an HTTP handler for the application.
type Handler struct {
	DB *DB
}

// Root renders the home page.
func (h *Handler) Root(w http.ResponseWriter, req *http.Request) {
	_ = h.DB.View(func(tx *Tx) error {
		return (&Template{}).Index(w, tx.Rooms())
	})
}

// Rooms renders a list of all rooms.
func (h *Handler) Rooms(w http.ResponseWriter, req *http.Request) {
	_ = h.DB.View(func(tx *Tx) error {
		return (&RoomTemplate{}).Index(w, tx.Rooms())
	})
}

// Room renders a single room.
func (h *Handler) Room(w http.ResponseWriter, req *http.Request) {
	_ = h.DB.View(func(tx *Tx) error {
		vars := mux.Vars(req)
		roomID, _ := strconv.Atoi(vars["id"])

		// Find room by ID.
		r := tx.Room(roomID)
		if r == nil {
			return errors.New("room not found")
		}

		return (&RoomTemplate{}).Show(w, r)
	})
}

// NewRoom renders the form for creating a room.
func (h *Handler) NewRoom(w http.ResponseWriter, req *http.Request) {
	_ = h.DB.View(func(tx *Tx) error {
		return (&RoomTemplate{}).New(w)
	})
}

// CreateRoom handles the creation of a new room.
func (h *Handler) CreateRoom(w http.ResponseWriter, req *http.Request) {
	r := &Room{
		Name: req.FormValue("name"),
	}

	// Create a new room.
	err := h.DB.Update(func(tx *Tx) error {
		id, _ := tx.Bucket([]byte("Rooms")).NextSequence()
		r.ID = int(id)
		return tx.SaveRoom(r)
	})
	if err != nil {
		(&Template{}).Error(w, err)
		return
	}

	// Redirect user to the room's URL.
	http.Redirect(w, req, fmt.Sprintf("/rooms/%d", r.ID), http.StatusFound)
}

// EditRoom renders the form for editing an existing room.
func (h *Handler) EditRoom(w http.ResponseWriter, req *http.Request) {
	_ = h.DB.View(func(tx *Tx) error {
		vars := mux.Vars(req)
		roomID, _ := strconv.Atoi(vars["id"])

		// Find room by ID.
		r := tx.Room(roomID)
		if r == nil {
			return errors.New("room not found")
		}

		return (&RoomTemplate{}).Edit(w, r)
	})
}

// UpdateRoom handles the update of an existing room.
func (h *Handler) UpdateRoom(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	roomID, _ := strconv.Atoi(vars["id"])

	// Update existing room.
	err := h.DB.Update(func(tx *Tx) error {
		// Find room by ID.
		r := tx.Room(roomID)
		if r == nil {
			return errors.New("room not found")
		}

		// Update properties and save room.
		r.Name = req.FormValue("name")
		return tx.SaveRoom(r)
	})
	if err != nil {
		(&Template{}).Error(w, err)
		return
	}

	// Redirect user to the room's URL.
	http.Redirect(w, req, fmt.Sprintf("/rooms/%d", roomID), http.StatusFound)
}

// DeleteRoom handles the deletion of a room.
func (h *Handler) DeleteRoom(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	roomID, _ := strconv.Atoi(vars["id"])

	// Delete the room.
	err := h.DB.Update(func(tx *Tx) error {
		return tx.DeleteRoom(roomID)
	})
	if err != nil {
		(&Template{}).Error(w, err)
		return
	}

	// Redirect user to the room listing.
	http.Redirect(w, req, "/rooms", http.StatusFound)
}

type Template struct{}

type RoomTemplate struct {
	Template
}
