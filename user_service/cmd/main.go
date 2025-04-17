package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"gitlab.com/e-market724/back-end/user_service/genproto/user_service"
	"gitlab.com/e-market724/back-end/user_service/pkg/db"
	"gitlab.com/e-market724/back-end/user_service/services"
	"gitlab.com/e-market724/back-end/user_service/storage"
	"google.golang.org/grpc"
)

func main() {

	db, err := db.ConnDB(context.Background())

	if err != nil {

		log.Println(err)
		return
	}

	log.Println(db)

	stotage := storage.NewUserRepo(db)

	log.Println(stotage)

	service := services.NewUserService(stotage)

	listen, err := net.Listen("tcp", "localhost:8080")

	if err != nil {

		log.Println(err)
		return
	}

	server := grpc.NewServer()

	user_service.RegisterUserServiceServer(server, service)

	fmt.Println("server serve on :8080")

	if err = server.Serve(listen); err != nil {
		log.Println(err.Error())
		return

	}

}
