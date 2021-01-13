package main

import (
	"github.com/joho/godotenv"
	"grpc-exaple/db"
	"grpc-exaple/grpc"
	"grpc-exaple/kafka"
	"log"
)

func main()  {

	//Load the .env file
	err := godotenv.Load("./.env_rename_me")
	if err != nil {
		log.Fatalf("Error loading .env file, please create one in the root directory: %s", err)
	}
	// connect db
	db.InitDB()
	// connect redis
	db.InitRedis("1")
	// init grpc server
	go grpc.InitGrpc()

	kafka.Consume()

}