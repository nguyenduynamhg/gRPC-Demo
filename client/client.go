package main

import (
	"context"
	"golang/golang/calculator/calculatorpb"
	"log"

	"google.golang.org/grpc"
)
func main() {
	//Client connection
	cc, err := grpc.Dial("localhost:50069",grpc.WithInsecure())  
	
	if err != nil{
		log.Fatalf("err while create listen %v", err)
	}
	//Đóng connection khi chạy xong
	defer cc.Close() 
	
	// Tạo client
	client := calculatorpb.NewCalculatorServiceClient(cc) 
	
	log.Printf("Server client: %f",client)
	callSum(client)
} 
func callSum(c calculatorpb.CalculatorServiceClient)  {
	log.Printf("calling sum api")
	resp, err := c.Sum(context.Background(), &calculatorpb.SumRequest{
		Num1: 5,
		Num2: 6,
	})
	if err != nil {
	  log.Fatalf("call sum api err %v", err)
	}
	log.Printf("call api response %v\n: ", resp.GetResult())
}