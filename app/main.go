package main

import (
	"net"
	"os"

	"xxx/logger"

	xc "xxx/app/xxx/controller"
	xg "xxx/app/xxx/gateway"
	xu "xxx/app/xxx/usecase"
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

	xxxUsecase := xu.NewXxxUsecase(xxxGateway)

	xxxController := xc.NewXxxController(xxxUsecase)

	pbx.RegisterXxxServiceServer(server, xxxController)

	logger.StartLog()

	server.Serve(listenPort)
}
