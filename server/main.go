package main

import (
	"context"
	pb "github.com/RianIhsan/go-learn-grpc/users"
	"sync"
)

type dataUserSever struct {
	pb.UnimplementedDataUsersServer
	mu    sync.Mutex
	users []*pb.User
}

func (d *dataUserSever) FindUserByEmail(ctx context.Context, user *pb.User) (*pb.User, error) {
	for _, v := range d.users {
		if v.Email == user.Email {
			return v, nil
		}
	}
	return nil, nil
}
