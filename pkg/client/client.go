package grpc

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/SXUOJ/backend/pb"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func Ping() {
	conn, err := grpc.Dial(getAddr(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewJudgerClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Ping(ctx, &pb.PingRequest{})

	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	fmt.Println(r)
}

func Judge(req *pb.JudgeRequest) {
	conn, err := grpc.Dial(getAddr(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewJudgerClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Judge(ctx, req)

	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	fmt.Println(r)
}

func getAddr() string {
	return viper.GetString("grpc.addr")
}
