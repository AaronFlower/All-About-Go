package main

import (
	"context"
	"log"
	"net"

	pb "github.com/aaronflower/dzone-stu/service.student/proto/student"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type iClass interface {
	Create(*pb.Student) []*pb.Student
	GetAll() []*pb.Student
}

type class struct {
	students []*pb.Student
}

func (c *class) Create(student *pb.Student) []*pb.Student {
	c.students = append(c.students, student)
	return c.students
}

func (c *class) GetAll() []*pb.Student {
	return c.students
}

type service struct {
	class iClass
}

func (s *service) CreateStudent(ctx context.Context, student *pb.Student) (*pb.Response, error) {
	s.class.Create(student)
	return &pb.Response{Created: true}, nil
}

func (s *service) GetAll(ctx context.Context, req *pb.GetRequest) (*pb.Response, error) {
	return &pb.Response{Students: s.class.GetAll()}, nil
}

func main() {
	var port = ":50051"
	cls := &class{}

	server := grpc.NewServer()
	pb.RegisterStudentServiceServer(server, &service{class: cls})
	reflection.Register(server)

	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err)
	}
	log.Fatal(server.Serve(listener))

}
