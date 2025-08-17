package main

import (
	"database/sql"
	grpcc "grpc_test/internal/delivery/grpc"
	"grpc_test/internal/delivery/grpc/pb"
	"grpc_test/internal/infrastructure/postgres"
	"grpc_test/internal/usecase"
	"time"

	_ "github.com/lib/pq"

	"fmt"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
)

func main() {
	dns := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("HOST"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
	)

	db, err := sql.Open("postgres", dns)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	for i := 0; i < 10; i++ {
		err := db.Ping()
		if err == nil {
			break
		}
		log.Printf("Waiting for database, attempt %d...", i+1)
		time.Sleep(2 * time.Second)
	}

	repo := postgres.NewPostgresRepo(db)
	usecase := usecase.NewUseCase(repo)
	service := grpcc.NewServer(usecase)

	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Failed to listen on port 8080:%v", err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterTodoServiceServer(grpcServer, service)

	log.Println("Server started on :8080")

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
