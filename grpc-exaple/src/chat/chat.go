package chat

import (
	"context"
	"grpc-exaple/db"
	"log"
)

type Server struct {
}

type Chat struct {
	id int64
	code string
	data string
}

func (s *Server) SayHello(ctx context.Context, message *MessageRequest) (*MessageResp, error) {
	log.Printf("get data redis: %s", db.GetRedis().Get(ctx,"test"))
	log.Printf("Receive message body from client: %s", message.Code)
	db.GetDB().QueryRow("INSERT INTO  chat(code,`data`) VALUES('aassasa','ssssssss')")
	return &MessageResp{Data: "Hello From the Server!",Code: "222", Id: 1222}, nil
}