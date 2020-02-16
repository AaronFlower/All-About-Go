package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/aaronflower/dzone-stu/service.student/proto/student"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	client := pb.NewStudentServiceClient(conn)

	res, err := client.CreateStudent(context.Background(), &pb.Student{
		Name: "Aaron",
		Age:  15,
	})

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("res = %+v\n", res)

	all, err := client.GetAll(context.Background(), &pb.GetRequest{})
	if err != nil {
		log.Fatal(err)
	}
	for _, v := range all.Students {
		fmt.Printf("v = %+v\n", v)
	}

}
