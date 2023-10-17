package main

import (
	"database/sql"
	"net"

	"github.com/feliperromao/go-grpc/internal/database"
	"github.com/feliperromao/go-grpc/internal/pb"
	"github.com/feliperromao/go-grpc/internal/service"
	_ "github.com/mattn/go-sqlite3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	db, err := sql.Open("sqlite3", "./db.sqlite")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	categoryDB := database.NewCategory(db)
	categorySrv := service.NewCategoryService(*categoryDB)

	grpcServer := grpc.NewServer()
	pb.RegisterCategoryServiceServer(grpcServer, categorySrv)
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}
	if err := grpcServer.Serve(listener); err != nil {
		panic(err)
	}
}
