package main

import (
	"net"
	"os"

	"clean/logger"

	hc "clean/app/hoge/controller"
	hg "clean/app/hoge/gateway"
	hu "clean/app/hoge/usecase"
	pbh "clean/proto/hoge"

	fc "clean/app/fuga/controller"
	fg "clean/app/fuga/gateway"
	fu "clean/app/fuga/usecase"
	pbf "clean/proto/fuga"

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

	hogeGateway := hg.NewHogeGateway()

	hogeUsecase := hu.NewHogeUsecase(hogeGateway)

	hogeController := hc.NewHogeController(hogeUsecase)

	pbh.RegisterHogeServiceServer(server, hogeController)

	fugaGateway := fg.NewFugaGateway()

	fugaUsecase := fu.NewFugaUsecase(fugaGateway)

	fugaController := fc.NewFugaController(fugaUsecase)

	pbf.RegisterFugaServiceServer(server, fugaController)

	logger.StartLog()

	server.Serve(listenPort)
}
