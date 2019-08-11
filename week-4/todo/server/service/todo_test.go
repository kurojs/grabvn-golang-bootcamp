package service

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xuanit/testing/todo/pb"
	"github.com/xuanit/testing/todo/server/repository/mocks"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

func TestGetToDo(t *testing.T) {
	mockToDoRep := &mocks.ToDo{}
	toDo := &pb.Todo{}
	req := &pb.GetTodoRequest{Id: "123"}
	mockToDoRep.On("Get", req.Id).Return(toDo, nil)
	service := ToDo{ToDoRepo: mockToDoRep}

	res, err := service.GetTodo(nil, req)

	expectedRes := &pb.GetTodoResponse{Item: toDo}

	assert.Nil(t, err)
	assert.Equal(t, expectedRes, res)
	mockToDoRep.AssertExpectations(t)
}

func TestCreateTodoOK(t *testing.T) {
	toDo := &pb.Todo{Id: "TestingID", Title: "Title", Completed: true, Description: "Non-description"}
	req := &pb.CreateTodoRequest{Item: toDo}
	mockTodoRep := &mocks.ToDo{}
	service := ToDo{ToDoRepo: mockTodoRep}

	mockTodoRep.On("Insert", req.Item).Return(nil)
	res, err := service.CreateTodo(context.Background(), req)
	expectedRes := &pb.CreateTodoResponse{Id: toDo.Id}

	assert.Nil(t, err)
	assert.Equal(t, expectedRes, res)
	mockTodoRep.AssertExpectations(t)
}

func TestCreateTodoFailed(t *testing.T) {
	existedToDo := &pb.Todo{Id: "TestingID", Title: "Title", Completed: true, Description: "Non-description"}
	errReq := &pb.CreateTodoRequest{Item: existedToDo}
	mockTodoRep := &mocks.ToDo{}
	service := ToDo{ToDoRepo: mockTodoRep}

	exceptedErr := grpc.Errorf(codes.Internal, "Could not insert item into the database")
	mockTodoRep.On("Insert", errReq.Item).Return(exceptedErr)

	res, err := service.CreateTodo(context.Background(), errReq)

	assert.NotNil(t, err)
	assert.Nil(t, res)
	mockTodoRep.AssertExpectations(t)
}

func TestListTodoOK(t *testing.T) {
	mockTodoRep := &mocks.ToDo{}
	service := ToDo{ToDoRepo: mockTodoRep}

	toDoList := []*pb.Todo{
		&pb.Todo{Id: "1", Completed: true},
		&pb.Todo{Id: "2", Completed: true},
		&pb.Todo{Id: "3", Completed: true},
	}
	req := &pb.ListTodoRequest{Limit: 3, NotCompleted: true}
	exceptedRes := &pb.ListTodoResponse{Items: toDoList}
	mockTodoRep.On("List", req.Limit, req.NotCompleted).Return(toDoList, nil)

	res, err := service.ListTodo(context.Background(), req)

	assert.Equal(t, res, exceptedRes)
	assert.LessOrEqual(t, len(res.Items), len(toDoList))
	assert.Nil(t, err)
	mockTodoRep.AssertExpectations(t)
}

func TestListTodoFailed(t *testing.T) {
	mockTodoRep := &mocks.ToDo{}
	service := ToDo{ToDoRepo: mockTodoRep}

	req := &pb.ListTodoRequest{Limit: 3, NotCompleted: true}
	exceptedErr := grpc.Errorf(codes.Internal, "Could not delete item from the database")
	mockTodoRep.On("List", req.Limit, req.NotCompleted).Return(nil, exceptedErr)

	res, err := service.ListTodo(context.Background(), req)

	assert.NotNil(t, err)
	assert.Nil(t, res)
	mockTodoRep.AssertExpectations(t)
}

func TestDeleteTodoOK(t *testing.T) {
	mockTodoRep := &mocks.ToDo{}
	service := ToDo{ToDoRepo: mockTodoRep}

	req := &pb.DeleteTodoRequest{Id: "1"}
	exceptedRes := &pb.DeleteTodoResponse{}
	mockTodoRep.On("Delete", req.Id).Return(nil)

	res, err := service.DeleteTodo(context.Background(), req)

	assert.Equal(t, res, exceptedRes)
	assert.Nil(t, err)
	mockTodoRep.AssertExpectations(t)
}

func TestDeleteTodoFail(t *testing.T) {
	mockTodoRep := &mocks.ToDo{}
	service := ToDo{ToDoRepo: mockTodoRep}

	req := &pb.DeleteTodoRequest{Id: "1"}
	exceptedErr := grpc.Errorf(codes.Internal, "Could not delete item from the database")

	mockTodoRep.On("Delete", req.Id).Return(exceptedErr)

	res, err := service.DeleteTodo(context.Background(), req)

	assert.Nil(t, res)
	assert.NotNil(t, err)
	mockTodoRep.AssertExpectations(t)
}
