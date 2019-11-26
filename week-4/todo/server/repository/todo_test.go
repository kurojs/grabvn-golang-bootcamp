// +build integration persistence

package repository

import (
	"testing"
	"time"

	"github.com/go-pg/pg"
	"github.com/stretchr/testify/suite"
	"github.com/xuanit/testing/todo/pb"
)

type ToDoRepositorySuite struct {
	db *pg.DB
	suite.Suite
	todoRep ToDoImpl
}

func (s *ToDoRepositorySuite) SetupSuite() {
	// Connect to PostgresQL
	s.db = pg.Connect(&pg.Options{
		User:                  "postgres",
		Password:              "example",
		Database:              "todo",
		Addr:                  "localhost" + ":" + "5433",
		RetryStatementTimeout: true,
		MaxRetries:            4,
		MinRetryBackoff:       250 * time.Millisecond,
	})

	// Create Table
	s.db.CreateTable(&pb.Todo{}, nil)

	s.todoRep = ToDoImpl{DB: s.db}
}

func (s *ToDoRepositorySuite) TearDownSuite() {
	s.db.DropTable(&pb.Todo{}, nil)
	s.db.Close()
}

func (s *ToDoRepositorySuite) TestInsert() {
	item := &pb.Todo{Id: "new_item", Title: "meeting"}
	err := s.todoRep.Insert(item)

	s.Nil(err)

	newTodo, err := s.todoRep.Get(item.Id)
	s.Nil(err)
	s.Equal(item, newTodo)
}

func (s *ToDoRepositorySuite) TestGet() {
	item := &pb.Todo{Id: "new_get_item", Title: "Test get"}
	s.todoRep.Insert(item)

	getToDo, err := s.todoRep.Get(item.Id)
	s.Nil(err)
	s.Equal(item, getToDo)
}

func (s *ToDoRepositorySuite) TestList() {
	items := []*pb.Todo{
		&pb.Todo{Id: "1", Title: "Test List 1", Completed: true},
		&pb.Todo{Id: "2", Title: "Test List 2", Completed: true},
		&pb.Todo{Id: "3", Title: "Test List 3", Completed: true},
	}
	err := s.todoRep.Insert(items[0])
	err := s.todoRep.Insert(items[1])
	err := s.todoRep.Insert(items[2])

	listToDo, err := s.todoRep.List(3, false)

	s.Nil(err)
	s.Equal(listToDo, items)
}

func (s *ToDoRepositorySuite) TestDelete() {
	item := &pb.Todo{Id: "new_delete_item", Title: "Test delete"}
	s.todoRep.Insert(item)

	err := s.todoRep.Delete(item.Id)
	s.Nil(err)
}

func TestToDoRepository(t *testing.T) {
	suite.Run(t, new(ToDoRepositorySuite))
}
