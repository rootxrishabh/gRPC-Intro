package main

import (
	"context"
	"log"

	svc "github.com/rootxrishabh/gRPC-Intro/client/client"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %s", err)
	}
	defer conn.Close()
	client := svc.NewInboxClient(conn)
	emails, err := client.GetEmails(context.Background(), &svc.EmailRequest{})
	if err != nil {
		log.Fatalf("could not get emails: %s", err)
	}
	log.Println(emails)
}
