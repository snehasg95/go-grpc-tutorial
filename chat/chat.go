package chat

import (
	"log"
	"golang.org/x/net/context"
)

type Server struct {
	}


func (s *Server) SayTestName(ctx context.Context, input *Message) (*Message, error) {
	log.Printf("----- Received message body from client----: %s", input.Body)
	return &Message{Body: "chat From the Server!"}, nil
