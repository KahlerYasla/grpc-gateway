package chat

import (
	"log"

	"/src/internal/service/chat/proto/gen"
)

type ChatService struct {
	gen.UnimplementedChatServiceServer
}

func (s *ChatService) StreamMessages(stream gen.ChatService_StreamMessagesServer) error {
	for {
		msg, err := stream.Recv()
		if err != nil {
			return err
		}

		log.Printf("Received message from %s: %s", msg.User, msg.Message)
		if err := stream.Send(msg); err != nil {
			return err
		}
	}
}
