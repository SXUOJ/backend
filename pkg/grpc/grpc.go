package grpc

import (
	"context"
	"log"
	"time"

	"github.com/SXUOJ/backend/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func Ping(addr string) (*pb.PongReply, error) {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
		return nil, err
	}
	defer conn.Close()

	c := pb.NewJudgerClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return c.Ping(ctx, &pb.PingRequest{})
}

func Judge(addr string, req *pb.JudgeRequest) (*pb.JudgeReply, error) {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
		return nil, err
	}
	defer conn.Close()

	c := pb.NewJudgerClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return c.Judge(ctx, req)
}
