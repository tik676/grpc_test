package grpc

import (
	"context"
	"grpc_test/internal/delivery/grpc/pb"
	"grpc_test/internal/domain"
	"grpc_test/internal/usecase"
)

type Server struct {
	pb.UnimplementedTodoServiceServer
	useCase *usecase.UseCase
}

func NewServer(useCase *usecase.UseCase) *Server {
	return &Server{useCase: useCase}
}

func (s *Server) CreateTask(ctx context.Context, req *pb.CreateTaskRequest) (*pb.TaskResponse, error) {
	task := s.useCase.Create(req.Title, req.Description)

	return &pb.TaskResponse{
		Task: &pb.Task{
			Id:          task.ID,
			Title:       task.Title,
			Description: task.Description,
			Completed:   task.Completed},
	}, nil
}

func (s *Server) ListTasks(ctx context.Context, req *pb.ListTasksRequest) (*pb.ListTasksResponse, error) {
	tasks := s.useCase.GetList()

	var pbTasks []*pb.Task
	for _, t := range tasks.Tasks {
		pbTasks = append(pbTasks, &pb.Task{
			Id:          int64(t.ID),
			Title:       t.Title,
			Description: t.Description,
			Completed:   t.Completed,
		})
	}

	return &pb.ListTasksResponse{
		Tasks: pbTasks,
	}, nil
}

func (s *Server) EditTask(ctx context.Context, req *pb.EditTaskRequest) (*pb.EditTaskResponse, error) {
	t := req.GetTask()
	domainTask := domain.Task{
		ID:          t.GetId(),
		Title:       t.GetTitle(),
		Description: t.GetDescription(),
		Completed:   t.GetCompleted(),
	}
	task := s.useCase.Edit(domainTask)
	return &pb.EditTaskResponse{
		Task: &pb.Task{
			Id:          task.ID,
			Title:       task.Title,
			Description: task.Description,
			Completed:   task.Completed,
		},
	}, nil
}

func (s *Server) DeleteTask(ctx context.Context, req *pb.DeleteTaskRequest) (*pb.TaskResponse, error) {
	task := s.useCase.Delete(int(req.Id))

	return &pb.TaskResponse{
		Task: &pb.Task{
			Id:          task.ID,
			Title:       task.Title,
			Description: task.Description,
			Completed:   task.Completed,
		},
	}, nil
}
