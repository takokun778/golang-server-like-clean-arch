package main

import (
	"net"
	"os"

	"xxx/logger"

	"xxx/app/infra/ent"
	xc "xxx/app/xxx/controller"
	xg "xxx/app/xxx/gateway"
	xi "xxx/app/xxx/interactor"
	pbx "xxx/proto/xxx"

	"google.golang.org/grpc"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "19003"
	}

	listenPort, err := net.Listen("tcp", ":"+port)
	if err != nil {
		panic(err)
	}

	server := grpc.NewServer()

	db := ent.DatabaseConnect()

	xxxGateway := xg.NewXxxGateway(db)

	xxxInteractor := xi.NewXxxInteractor(xxxGateway)

	xxxController := xc.NewXxxController(xxxInteractor)

	pbx.RegisterXxxServiceServer(server, xxxController)

	logger.StartLog()

	server.Serve(listenPort)
}
