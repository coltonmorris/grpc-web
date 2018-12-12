// Based off of the file at: https://github.com/improbable-eng/grpc-web/blob/master/example/go/exampleserver/exampleserver.go
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	// pb "git.tcncloud.net/m/protoc-gen-state/example/dogApp/backend/proto-types"
	pb "./proto-types"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"

	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"golang.org/x/net/context"
)

const (
	port = ":9090"
)

type ReadingListServer struct{}

// type ReadingListServer struct {
//   bookList []pb.Book
// }

func (s *ReadingListServer) AddBook(ctx context.Context, in *pb.Book) (sm *pb.ServerMessage, err error) {
	fmt.Println("in AddBook")
	return &pb.ServerMessage{Message: "Added book"}, nil
}

// func (s *ReadingListServer) GetReadBooks(ctx context.Context, in *pb.SpecialRequest) (sm *pb.ServerMessage, err error) {
//   fmt.Println("In GetReadBooks")
//   return &pb.ServerMessage{Message: "It's me, mario"},nil
// }

func (s *ReadingListServer) RemoveBook(ctx context.Context, in *pb.Book) (sm *pb.ServerMessage, err error) {
	fmt.Println("In RemoveBook")
	return &pb.ServerMessage{Message: "Removed book"}, nil
}

func main() {
	s := grpc.NewServer()
	pb.RegisterReadingListServer(s, &ReadingListServer{})

	grpclog.SetLogger(log.New(os.Stdout, "go server: ", log.LstdFlags))

	wrappedServer := grpcweb.WrapServer(s)
	handler := func(resp http.ResponseWriter, req *http.Request) {
		wrappedServer.ServeHTTP(resp, req)
	}

	httpServer := http.Server{
		Addr:    port,
		Handler: http.HandlerFunc(handler),
	}

	grpclog.Printf("Starting server at localhost%v", port)

	if err := httpServer.ListenAndServe(); err != nil {
		grpclog.Fatalf("failed starting http server: %v", err)
	}
}
