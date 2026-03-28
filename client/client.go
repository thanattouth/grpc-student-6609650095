package main

import (
	"context"
	"log"
	"time"

	pb "grpc-student/studentpb"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewStudentServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.GetStudent(ctx, &pb.StudentRequest{Id: 101})
	if err == nil {
		log.Printf("GetStudent -> Name: %s, Phone: %s", res.Name, res.Phone)
	}

	listRes, err := client.ListStudents(ctx, &pb.Empty{})
	if err != nil {
		log.Fatalf("Error calling ListStudents: %v", err)
	}

	log.Println("Student List:")
	for _, s := range listRes.GetStudent() {
		log.Printf("- ID: %d, Name: %s, Phone: %s", s.Id, s.Name, s.Phone)
	}
}
