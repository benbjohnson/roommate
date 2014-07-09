package main

import (
	"flag"
	"log"
	"net/http"
	"os/user"
	"path/filepath"

	"github.com/benbjohnson/roommate"
)

func main() {
	path := flag.String("d", "", "data file")
	addr := flag.String("addr", ":50000", "bind address")
	flag.Parse()
	log.SetFlags(0)

	// Default the path to a file in the user's home directory.
	if *path == "" {
		u, err := user.Current()
		if err != nil {
			log.Fatal(err)
		}
		*path = filepath.Join(u.HomeDir, ".roommate")
	}

	// Open the database.
	var db roommate.DB
	if err := db.Open(*path, 0600); err != nil {
		log.Fatal("db:", err)
	}

	// Start HTTP server.
	log.Printf("listening on http://localhost%s", *addr)
	log.Fatal(http.ListenAndServe(*addr, roommate.NewHandler(&db)))
}
