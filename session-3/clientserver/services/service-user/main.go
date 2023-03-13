package main

import (
	"context"
	"log"

	"GLIM_Hacktiv8/golang-intermediate/session-3/clientserver/common/model"

	"github.com/golang/protobuf/ptypes/empty"
)

var localStorage *model.UserList

func init() {
	localStorage = new(model.UserList)
	localStorage.List = make([]*model.User, 0)
}

type UsersServer struct {
	model.UnimplementedUsersServer
}

func (u UsersServer) Register(ctx context.Context, param *model.User) (*empty.Empty, error) {
	log.Printf("Register user request %+v\n", param)
	localStorage.List = append(localStorage.List, param)

	log.Println("Registering user", param.String())

	return new(empty.Empty), nil
}
