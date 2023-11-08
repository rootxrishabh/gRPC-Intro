package main

import (
	"context"
	"log"
	"net"

	"github.com/rootxrishabh/gRPC-Intro/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	server.UnimplementedInboxServer
}

func (s *Server) GetEmails(ctx context.Context, er *server.EmailRequest) (*server.EmailResponse, error) {
	return &server.EmailResponse{
		Emails: getSampleEmails(),
	}, nil
}

func main() {
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()
	reflection.Register(s)
	server.RegisterInboxServer(s, &Server{})
	log.Println("Server started at port 8080")
	if err := s.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func getSampleEmails() []*server.Email {
	smapleEmail := []*server.Email{
		{
			From:    "admin@mycompany.com",
			To:      "employee@mycompany.com",
			Subject: "Letter of Acceptance",
			Body:    "Congratulations! We are pleased to offer you the position of Software Engineer at MyCompany. We feel that your skills and background will be valuable assets to our team. We are excited to have you join us in our mission to make the world a better place through software.",
		},
		{
			From:    "cto@mycompany.com",
			To:      "cep@mycompany.com",
			Subject: "Great 1st quator results!",
			Body:    "I am pleased to announce that we have exceeded our revenue targets for the first quarter. We are on track to have a record breaking year. Keep up the good work!",
		},
	}
	return smapleEmail
}
