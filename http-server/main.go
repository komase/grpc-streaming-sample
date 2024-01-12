package main

import (
	"context"
	"encoding/json"
	"errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "grpc-straming-sample/grpc-sever/gen/go/hello"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"
)

var (
	addr = "localhost:50051"
)

func getAndSendMessages(client pb.MessengerClient, w http.ResponseWriter, r *http.Request) error {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	name := r.FormValue("name")
	count, err := strconv.Atoi(r.FormValue("count"))
	if err != nil {
		return err
	}

	// gRPC request
	stream, err := client.SayHello(ctx, &pb.HelloRequest{Name: name, Count: int32(count)})
	if err != nil {
		return err
	}

	f, ok := w.(http.Flusher)
	if !ok {
		return errors.New("unexpected type of web server")
	}
	for {
		message, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		if err := json.NewEncoder(w).Encode(message); err != nil {
			return err
		}
		f.Flush()
	}
	return nil
}

func sendErrorMessage(w http.ResponseWriter) error {
	w.WriteHeader(http.StatusInternalServerError)
	response := "internal server error"
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Print("Error encoding JSON response:", err)
		return err
	}
	return nil
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// setup gRPC client
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.Dial(addr, opts...)
	if err != nil {
		log.Printf("failed to gprc connection: %v", err)
		if err := sendErrorMessage(w); err != nil {
			log.Fatal("failed to send error message: %v", err)
		}
		return
	}
	defer conn.Close()
	client := pb.NewMessengerClient(conn)

	if err := getAndSendMessages(client, w, r); err != nil {
		log.Printf("failed to getAndSendMessages %v", err)
		if err := sendErrorMessage(w); err != nil {
			log.Fatalf("failed to send error message: %v", err)
		}
		return
	}
}

func main() {
	http.HandleFunc("/v1/hello", helloHandler)
	log.Fatal(http.ListenAndServe(":8082", nil))
}
