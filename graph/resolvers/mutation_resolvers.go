package resolvers

import (
	"context"
	"trackingApp/graph/generated"
	"trackingApp/graph/models"
)

func (m *mutationResolver) Login(ctx context.Context, input models.LoginInput) (string, error) {
	return m.Services.Login(ctx, input)
}

func (m *mutationResolver) CreateUser(ctx context.Context, input models.CreateUserInput) (bool, error) {
	return m.Services.CreateUser(ctx, input)
}

func (m *mutationResolver) UpdateUser(ctx context.Context, input models.UpdateUserInput) (bool, error) {
	return m.Services.UpdateUser(ctx, input)
}

func (m *mutationResolver) UpdateAdmin(ctx context.Context, input models.UpdateAdminInput) (bool, error) {
	return m.Services.UpdateAdmin(ctx, input)
}

func (m *mutationResolver) DeleteUser(ctx context.Context, userID string) (bool, error) {
	return m.Services.DeleteUser(ctx, userID)
}

func (m *mutationResolver) ActivateUser(ctx context.Context, userID string) (bool, error) {
	return m.Services.ActivateUser(ctx, userID)
}

func (m *mutationResolver) ChangePassword(ctx context.Context, input models.ChangePasswordInput) (bool, error) {
	return m.Services.ChangePassword(ctx, input)
}

func (m *mutationResolver) CreateItem(ctx context.Context, input models.CreateItemInput) (bool, error) {
	return m.Services.CreateItem(ctx, input)
}

func (m *mutationResolver) UpdateItem(ctx context.Context, input models.UpdateItemInput) (bool, error) {
	return m.Services.UpdateItem(ctx, input)
}

func (m *mutationResolver) DeleteItem(ctx context.Context, itemID string) (bool, error) {
	return m.Services.DeleteItem(ctx, itemID)
}

func (m *mutationResolver) LocationGen(ctx context.Context, input models.LocationGenInput) (bool, error) {
	return m.Services.LocationGen(ctx, input)
}

func (m *mutationResolver) CreateRent(ctx context.Context, input models.CreateRentInput) (bool, error) {
	return m.Services.CreateRent(ctx, input)
}

func (m *mutationResolver) DeleteRent(ctx context.Context, rentID string) (bool, error) {
	return m.Services.DeleteRent(ctx, rentID)
}

func (m *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{m} }

type mutationResolver struct{ *Resolver }
