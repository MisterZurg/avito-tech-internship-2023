package main

import (
	"avito-tech-internship-2023/internal/config"
	"avito-tech-internship-2023/internal/repository"
	"avito-tech-internship-2023/internal/service"
	pb "avito-tech-internship-2023/proto"
	"context"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
	"os"
)

const (
	DBType = "postgres"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		fmt.Println("Cannot get Config")
		return
	}

	// TODO db connect
	pgDSN := config.GetPostgresConnectionString(DBType, cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)
	pgDB, err := repository.NewRepository(pgDSN)
	if err != nil {
		panic("error parsing config")
	} else {
		fmt.Println("Connected tho")
	}

	restStuff(pgDB)
}

func restStuff(pgDB *repository.Repository) {
	// Create a listener on TCP port
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	// postgres.InitTables(pgDB.DB)

	// srv := service.NewService(pgDB)

	// Create a gRPC server object
	s := grpc.NewServer()

	// Attach the services to the server
	// Attach the services & db to the server
	service.NewService(s, pgDB)

	// Serve gRPC server
	log.Println("Serving gRPC on 0.0.0.0:8080")
	go func() {
		log.Fatalln(s.Serve(lis))
	}()

	maxMsgSize := 1024 * 1024 * 20
	// Create a client connection to the gRPC server we just started
	// This is where the gRPC-Gateway proxies the requests
	conn, err := grpc.DialContext(
		context.Background(),
		"0.0.0.0:8080",
		grpc.WithBlock(),
		grpc.WithInsecure(),
		grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(maxMsgSize), grpc.MaxCallSendMsgSize(maxMsgSize)),
	)
	if err != nil {
		log.Fatalln("Failed to dial server:", err)
	}

	gwmux := runtime.NewServeMux()
	// Register User Service
	err = pb.RegisterUsersServiceHandler(context.Background(), gwmux, conn)
	if err != nil {
		log.Fatalln("Failed to register gateway:", err)
	}
	newServer := pb.RegisterSegmentsServiceHandler(context.Background(), gwmux, conn)
	if newServer != nil {
		log.Fatalln("Failed to register gateway:", newServer)
	}

	gwServer := &http.Server{
		Addr:    fmt.Sprintf(":%s", "8090"), // os.Getenv("server_port"))
		Handler: gwmux,
	}

	log.Println(fmt.Sprintf("Serving gRPC-Gateway on %s:%s", os.Getenv("server_host"), os.Getenv("server_port")))
	log.Fatalln(gwServer.ListenAndServe())
}
