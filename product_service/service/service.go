package service

import (
	"gitlab.com/e-market724/back-end/product_service/genproto/product_service"
	"gitlab.com/e-market724/back-end/product_service/storage"
	"log"
	"net"

	"google.golang.org/grpc"
)

func Service() {

	listen, err := net.Listen("tcp", ":8000")

	if err != nil {
		log.Println(err)
		return
	}

	server := grpc.NewServer()

	storage := storage.NewStorage()

	productService := NewProductService(storage)

	product_service.RegisterPingServer(server, productService)

	server.Serve(listen)
}
