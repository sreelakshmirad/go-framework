package core

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// Config :
// DBConnString prepared by a mysql helper
// to be includ
type Config struct {
	DBConnString string
}

// Server :
// db methods and db config are the members
type Server struct {
	*Config
}

// Start a new http server
// and listen and serve to the defined port with
// defined routes
func (s *Server) Start() {

	ser := &http.Server{
		Addr:           ":8080",
		Handler:        s.InitRouter(),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	fmt.Println("Inside the start")
	log.Fatal(ser.ListenAndServe())
}

// NewServer :
// Create a new server with
// DB Configurations
// @param : c Config is pointer to Config object
// @return : Server; new server
func NewServer(c *Config) *Server {
	s := &Server{
		Config: c,
	}

	// db := helpers.MustPrepareDB(c.DBConnString)

	// s.db = &data.DB{db}
	return s
}
