package main

import (
	"log"
	"net"

	"github.com/Lambels/gRpc-Todo-App/service"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc"
)

func main() {
	db, err := sqlx.Open("mysql", "root:@tcp(localhost:3306)/todo")
	if err != nil {
	 	log.Fatal(err)
	}

	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	gs := grpc.NewServer()
	ts := NewTodoServer(db)

	pb.RegisterTasksServiceServer(gs, ts)

	log.Fatal(gs.Serve(l))
}