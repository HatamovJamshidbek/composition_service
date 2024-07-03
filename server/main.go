package main

import (
	"composition_service/config"
	pb "composition_service/genproto"
	"composition_service/service"
	"composition_service/storage/postgres"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	cnf := config.Config{}
	db, err := postgres.ConnectionDb(&cnf)
	if err != nil {
		log.Fatalf("error:->%s", err.Error())
	}
	listen, err := net.Listen("tcp", ":8087")
	if err != nil {
		log.Fatalf("error:->%s", err.Error())
	}
	grpcServer := grpc.NewServer()
	pb.RegisterCompositionServiceServer(grpcServer, service.NewCompositionService(postgres.NewCompositionRepository(db), postgres.NewTrackRepository(db)))
	log.Printf("Listen:%d", listen.Addr())
	err = grpcServer.Serve(listen)
	if err != nil {
		log.Fatalf("error:->%s", err.Error())
	}
}
