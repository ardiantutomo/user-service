package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"user-service/graph/generated"
	"user-service/graph/model"

	"golang.org/x/crypto/bcrypt"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input model.UserInput) (*model.User, error) {
	user := model.User{}
	user.Fullname = input.Fullname
	user.Email = input.Email
	hash, _ := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	user.Password = string(hash)
	res, err := r.userService.Save(user)
	return &res, err
}

func (r *queryResolver) GetAllUsers(ctx context.Context) ([]*model.User, error) {
	res, err := r.userService.GetAll()
	r.users = []*model.User{}
	for _, data := range res {
		r.users = append(r.users, &data)
	}
	return r.users, err
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
