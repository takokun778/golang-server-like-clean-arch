package main

import (
	"net"
	"os"

	"xxx/logger"

	"xxx/app/infra/ent"
	xp "xxx/app/infra/proto/xxx"
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

	listener, err := net.Listen("tcp", ":"+port)

	if err != nil {
		panic(err)
	}

	server := grpc.NewServer()

	db := ent.DatabaseConnect()

	xxxGateway := xg.NewXxxGateway(db)

	xxxAdapter := xp.NewXxxAdapter()

	xxxInteractor := xi.NewXxxInteractor(xxxGateway, xxxAdapter)

	xxxController := xc.NewXxxController(xxxInteractor)

	xxxService := xp.NewXxxService(xxxController)

	pbx.RegisterXxxServiceServer(server, xxxService)

	logger.StartLog()

	server.Serve(listener)
}
