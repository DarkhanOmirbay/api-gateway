package main

import (
	"context"
	pv1 "github.com/DarkhanOmirbay/proto/proto/gen/go/post"
	ssov1 "github.com/DarkhanOmirbay/proto/proto/gen/go/sso"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net/http"
	"os"
	"os/signal"
)

func runRest() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := pv1.RegisterPostServiceHandlerFromEndpoint(ctx, mux, "localhost:50051", opts)
	if err != nil {
		panic(err)
	}
	err = ssov1.RegisterAuthHandlerFromEndpoint(ctx, mux, "localhost:44044", opts)
	log.Printf("server listening at 8081")
	if err := http.ListenAndServe(":8081", mux); err != nil {
		panic(err)
	}
}
func main() {
	go runRest()

	// Ожидание сигнала остановки (например, нажатия Ctrl+C)
	// для того, чтобы приложение не завершилось сразу после запуска.
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	<-stop
	log.Println("Server stopped")
}
