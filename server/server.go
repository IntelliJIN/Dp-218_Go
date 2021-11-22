package main

import (
	"2gis/pb"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const step = 0.1

func main() {

	svr := NewServer()
	svr.Run()
	grpcServer := grpc.NewServer()

	pb.RegisterScooterServiceServer(grpcServer, svr)

	listener, err := net.Listen("tcp", ":8000")
	if err != nil {
		panic(err)
	}
	go func() {
		fmt.Println("grpc server started: 8000")
		log.Fatal(grpcServer.Serve(listener))
	}()

	http.HandleFunc("/", GISHandler)
	http.HandleFunc("/scooter", svr.ScooterHandler)

	fmt.Println("http server started: 8000")
	http.ListenAndServe(":9000", nil)

}

func GISHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "../index.html")
}

func (ss *Server) ScooterHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("new client connected")
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	client := &Client{
		w:    w,
		done: make(chan struct{}),
	}
	ss.AddClient(client)

	<-client.done
	fmt.Println("connection closed")

}

type Client struct {
	w    io.Writer
	done chan struct{}
}

type Server struct {
	client map[int]*Client
	taken  map[int]bool
	codes  map[int]int
	in     chan *pb.ClientMessage
}

func NewServer() *Server {
	return &Server{
		client: make(map[int]*Client),
		taken:  make(map[int]bool),
		codes:  make(map[int]int),
		in:     make(chan *pb.ClientMessage),
	}
}

func (ss *Server) AddClient(c *Client) {
	ss.client[1] = c
}

func (ss *Server) Register(msg *pb.ClientRequest, stream pb.ScooterService_RegisterServer) error {
	return nil
}

func (ss *Server) Receive(stream pb.ScooterService_ReceiveServer) error {
	var err error

	for {
		msg, err := stream.Recv()
		if err != nil {
			fmt.Println(err)
			err = status.Errorf(codes.Internal, "unexpected error %v", err)
			break
		}

		ss.in <- msg

	}

	return err
}

func (ss *Server) Run() {
	go func() {
		for {
			select {
			case msg := <-ss.in:
				// if !ok {
				// 	fmt.Printf("scooter %d stopped\n", msg.ID)
				// 	delete(ss.scooters, msg.ID)
				// }

				fmt.Println(msg)
				var buf bytes.Buffer
				json.NewEncoder(&buf).Encode(msg)

				for _, client := range ss.client {

					go func(c *Client) {
						if _, err := fmt.Fprintf(c.w, "data: %v\n\n", buf.String()); err != nil {
							fmt.Println(err)
							delete(ss.client, 1)
							close(c.done)
							return
						}

						if f, ok := c.w.(http.Flusher); ok {
							f.Flush()
						}
						fmt.Printf("data: %v\n", buf.String())
					}(client)

				}
			}
		}
	}()
}
