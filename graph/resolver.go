package graph

import (
	"go-graphql-clean/graph/model"
	"go-graphql-clean/service"
)

type Resolver struct {
	//masukin apa aja yg dibutuhin
	userService service.UserService
	users       []*model.User
}

func NewUserResolver(userService service.UserService) *Resolver {
	return &Resolver{
		userService: userService,
	}
}
