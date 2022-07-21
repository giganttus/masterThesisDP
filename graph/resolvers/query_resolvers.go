package resolvers

import (
	"context"
	"trackingApp/graph/generated"
	"trackingApp/graph/models"
)

func (q *queryResolver) UserSettings(ctx context.Context, userID string) (*models.UserSettings, error) {
	return q.Services.UserSettings(ctx, userID)
}
func (q *queryResolver) ProfileSettings(ctx context.Context) (*models.UserSettings, error) {
	return q.Services.ProfileSettings(ctx)
}

func (q *queryResolver) GetUsers(ctx context.Context) ([]*models.User, error) {
	return q.Services.GetUsers(ctx)
}

func (q *queryResolver) GetRoleNames(ctx context.Context) ([]*models.Role, error) {
	return q.Services.GetRoleNames(ctx)
}

func (q *queryResolver) GetItems(ctx context.Context) ([]*models.Item, error) {
	return q.Services.GetItems(ctx)
}

func (q *queryResolver) GetInputList(ctx context.Context) ([]*models.InputList, error) {
	return q.Services.GetInputList(ctx)
}

func (q *queryResolver) GetRents(ctx context.Context) ([]*models.Rent, error) {
	return q.Services.GetRents(ctx)
}

func (q *Resolver) Query() generated.QueryResolver { return &queryResolver{q} }

type queryResolver struct{ *Resolver }
