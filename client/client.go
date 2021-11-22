package main

import (
	"2gis/pb"
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc"
)

const step = 0.5
const interval = 1

func main() {
	conn, err := grpc.DialContext(context.Background(), ":8000", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	sclient := pb.NewScooterServiceClient(conn)
	stream, err := sclient.Receive(context.Background())
	if err != nil {
		panic(err)
	}

	cl := NewClient(1, 50.0, 40.0, stream)
	cl.Run(1)
}

type Client struct {
	Id uint64
	X  float64
	Y  float64
	// In  chan ServerMessage
	stream pb.ScooterService_ReceiveClient
}

// x, y are init coordinates of scooter
func NewClient(id uint64, x, y float64, stream pb.ScooterService_ReceiveClient) *Client {
	return &Client{
		Id: id,
		X:  x,
		Y:  y,
		// In:  in,
		stream: stream,
	}
}

func (s *Client) Run(interval int) {

	intPol := time.Duration(interval) * time.Second

	fmt.Println("executing run in client")
	// x, y := randomStep()

	for {

		//TODO change direction make it random

		s.X, s.Y = s.X+step, s.Y+step
		// send location to server
		msg := &pb.ClientMessage{
			Id: s.Id,
			X:  s.X,
			Y:  s.Y,
		}
		err := s.stream.Send(msg)
		if err != nil {
			panic(err)
		}
		fmt.Println("after send client")
		time.Sleep(intPol)
	}
}

func randomStep() (float64, float64) {
	return 0, 0
}
