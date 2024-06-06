package main

import (
	"database/sql"
	"log"
	"net"

	"github.com/sudiptob2/microgo/auth/internal/implementation/auth"
	pb "github.com/sudiptob2/microgo/auth/proto"
	"google.golang.org/grpc"
)

const (
	dbDriver = "mysql"
	dbUser   = "root"
	dbPass   = "Admin@123"
)

var db *sql.DB

func main() {
	var err error
	dsn := dbUser + ":" + dbPass + "@tcp(localhost:3306)/"
	db, err = sql.Open(dbDriver, dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			log.Printf("Error closing database connection: %v", err)
		}
	}()
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	// GRPC server setup step
	grpcServer := grpc.NewServer()
	authServer := auth.New(db)
	pb.RegisterAuthServiceServer(grpcServer, authServer)
	listener, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	log.Printf("Starting server on :%v\n", listener.Addr().String())
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

}
