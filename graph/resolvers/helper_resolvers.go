package resolvers

import (
	"context"
	"trackingApp/graph/generated"
	"trackingApp/graph/models"
	"trackingApp/middleware"
)

func (r *itemResolver) Rents(ctx context.Context, obj *models.Item) ([]*models.Rent, error) {
	return r.Services.GetRentsForItem(ctx, obj)
}

func (r *rentResolver) ItemsID(ctx context.Context, obj *models.Rent) (*models.Item, error) {
	return middleware.GetItemLoader(ctx).Load(obj.ItemsID)
}

// Item returns generated.ItemResolver implementation.
func (r *Resolver) Item() generated.ItemResolver { return &itemResolver{r} }

// Rent returns generated.RentResolver implementation.
func (r *Resolver) Rent() generated.RentResolver { return &rentResolver{r} }

type itemResolver struct{ *Resolver }

type rentResolver struct{ *Resolver }
