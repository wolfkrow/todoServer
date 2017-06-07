package main

import (
  "log"
  "net"
  "golang.org/x/net/context"
  "google.golang.org/grpc"
  pb "TODOLIST/lib"
)

const (
  port = ":8080"
)

type server struct{}

//func (s *server) Process(ctx context.Context, in *pb.Config) (*pb.Response, error) {
//  return &pb.Response{Message: "Hello " + in.Name}, nil
//}

func (s *server) CreateTask(ctx context.Context, in *pb.Task) (*pb.Response, error){
    log.Printf("Create object Task...%v %v", in.Title, in.Body)
    return &pb.Response{Message: "OK created" }, nil
}

func (s *server) GetTasks(ctx context.Context, in *pb.Options) (*pb.TaskList, error){
    return nil, nil
}

func (s *server)  DeleteTask(ctx context.Context, in *pb.Task) (*pb.Response, error){
        return &pb.Response{Message: "OK deleted" }, nil

}

func main() {
  lis, err := net.Listen("tcp", port)
  if err != nil {
    log.Fatalf("failed to listen: %v", err)
  }
  s := grpc.NewServer()
  log.Printf("Listening...")
  pb.RegisterRequesterServer(s, &server{})
  s.Serve(lis)
}

