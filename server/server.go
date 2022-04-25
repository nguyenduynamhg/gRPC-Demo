package main

import (
	"context"
	"fmt"
	"golang/golang/calculator/calculatorpb"
	"log"
	"net"

	"google.golang.org/grpc"
)
type server struct {}

func (*server) Sum(ctx context.Context, req *calculatorpb.SumRequest) (*calculatorpb.SumResponse, error) {
	log.Printf("Sum called...")
	resp := &calculatorpb.SumResponse{
		Result: req.GetNum1()+ req.GetNum2(),
	}
	return resp, nil
}
func main() {
	lis, err := net.Listen("tcp","0.0.0.0:50069") //tạo port
	if err != nil{
		log.Fatalf("err while create listen %v", err)
	}

	s := grpc.NewServer() // Tạo server

	calculatorpb.RegisterCalculatorServiceServer(s, &server{}) //Đăng kí server
	
	fmt.Println("calculator is running at 50069...")
	err = s.Serve(lis) // Run server

	if err != nil{
		log.Fatalf("err while create listen %v", err)
	}
} 