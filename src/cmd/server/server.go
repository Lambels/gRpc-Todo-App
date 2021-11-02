package main

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/Lambels/gRpc-Todo-App/service"
	"github.com/jmoiron/sqlx"
	"google.golang.org/protobuf/proto"
)

func NewTodoServer(repo *sqlx.DB) *todoServer {
	return &todoServer {
		repo: repo,
	}
}

type todoServer struct {
	repo *sqlx.DB

	pb.UnimplementedTasksServiceServer
}

func (s *todoServer) List(ctx context.Context, in *pb.ListRequest) (*pb.ListResponse, error) {
	var rws *sql.Rows
	var err error
	var out = new(pb.ListResponse)

	if proto.Equal(in, &pb.ListRequest{}) {
		rws, err = s.repo.QueryContext(ctx, "SELECT * FROM tasks ORDER BY id DESC LIMIT 7")
		if err != nil {
			return out, err
		}
	} else {
		rws, err = s.repo.QueryContext(ctx, "SELECT * FROM tasks WHERE message LIKE ? AND done = ? ORDER BY id DESC LIMIT 7", "%"+in.Message+"%", in.Done)
		if err != nil {
			return out, err
		}
	}

	defer rws.Close()

	var dbStruct = struct {
		Id int64 
		Message string
		Done bool
	} {}

	for rws.Next() {
		if err := rws.Scan(&dbStruct.Id, &dbStruct.Message, &dbStruct.Done); err != nil {
			return out, err
		}

		out.Tasks = append(out.Tasks, &pb.TodoTask {Id: dbStruct.Id, Message: dbStruct.Message, Done: dbStruct.Done})
	}

	if rws.Err() != nil {
		return out, rws.Err()
	}

	return out, nil
}

func (s *todoServer) Done(ctx context.Context, in *pb.TodoTask) (*pb.Void, error) {
	var void = new(pb.Void)

	if proto.Equal(in, new(pb.TodoTask)) {
		return void, fmt.Errorf("there wasnt a provided task")
	}

	_, err := s.repo.NamedExecContext(ctx, "UPDATE tasks SET done=1 WHERE id=:id", in)

	return void, err
}

func (s *todoServer) Add(ctx context.Context, in *pb.Message) (*pb.Void, error) {
	var void = new(pb.Void)

	if in.GetMessage() == "" {
		return void, fmt.Errorf("there wasnt provided a message")
	}

	_, err := s.repo.NamedExecContext(ctx, "INSERT INTO tasks(message, done) VALUES(:message, 0)", in)
	return void, err
}