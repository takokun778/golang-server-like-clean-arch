package main

import (
	"net"
	"os"

	"xxx/logger"

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

	xxxGateway := xg.NewXxxGateway()

	xxxUsecaseCreate := xi.NewXxxInteractorCreate(xxxGateway)
	xxxUsecaseRead := xi.NewXxxInteractorRead(xxxGateway)
	xxxUsecaseReadAll := xi.NewXxxInteractorReadAll(xxxGateway)
	xxxUsecaseUpdate := xi.NewXxxInteractorUpdate(xxxGateway)
	xxxUsecaseDelete := xi.NewXxxInteractorDelete(xxxGateway)

	xxxController := xc.NewXxxController(
		xxxUsecaseCreate,
		xxxUsecaseRead,
		xxxUsecaseReadAll,
		xxxUsecaseUpdate,
		xxxUsecaseDelete,
	)

	pbx.RegisterXxxServiceServer(server, xxxController)

	logger.StartLog()

	server.Serve(listenPort)
}
