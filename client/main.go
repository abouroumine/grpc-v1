package main

import (
	pb "abouroumine.com/grpc-v1/usermg"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"time"
)

const (
	address = "localhost:50051"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Did not Connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewUserMgClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	var newUsers = make(map[string]int32)
	newUsers["Alice"] = 43
	newUsers["Ayoub"] = 32
	for name, age := range newUsers {
		fmt.Println("Step 8")
		p := pb.NewUser{Name: name, Age: age}
		r, err := c.CreateUser(ctx, &p)
		if err != nil {
			log.Fatalf("Could not Create User: %v\n", err)
		}
		log.Printf("User Details:\nName: %s\nAge: %d\nID: %d\n", r.GetName(), r.GetAge(), r.GetUserid())
	}
}
