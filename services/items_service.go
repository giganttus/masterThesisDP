package services

import (
	"context"
	"errors"
	"math/rand"
	"trackingApp/graph/models"
	"trackingApp/middleware"
)

func (s *Services) CreateItem(ctx context.Context, input models.CreateItemInput) (bool, error) {
	res := s.UsersRepo.AllowAccess(ctx, "CRUD items")
	if res != true {
		return false, ErrForbidden
	}

	item := &models.Item{
		TypeID:   input.TypeID,
		Lon:      input.Lon,
		Lat:      input.Lat,
		BrokenID: input.BrokenID,
	}

	return s.ItemsRepo.CreateItem(item)
}

func (s *Services) GetItems(ctx context.Context) ([]*models.Item, error) {
	res := s.UsersRepo.AllowAccess(ctx, "CRUD items")
	if res != true {
		return nil, ErrForbidden
	}

	return s.ItemsRepo.GetItems()
}

func (s *Services) GetInputList(ctx context.Context) ([]*models.InputList, error) {
	res := s.UsersRepo.AllowAccess(ctx, "CRUD items")
	if res != true {
		return nil, ErrForbidden
	}

	typeNames, err := s.ItemsRepo.GetTypeNames()
	if err != nil {
		return nil, err
	}

	inputList := make([]*models.InputList, len(typeNames), 100)
	for i := range typeNames {
		inputList[i] = &models.InputList{Name: typeNames[i].Name}
	}

	inputTypes, errInTp := s.ItemsRepo.GetInputTypes()
	if errInTp != nil {
		return nil, errInTp
	}

	for i := range inputList {
		for j := range inputTypes {
			if inputList[i].Name == inputTypes[j].Name {
				atr := models.Attributes{ID: inputTypes[j].ID, Value: inputTypes[j].Title}
				inputList[i].Attributes = append(inputList[i].Attributes, &atr)
			}
		}
	}

	return inputList, nil
}

func (s *Services) UpdateItem(ctx context.Context, input models.UpdateItemInput) (bool, error) {
	res := s.UsersRepo.AllowAccess(ctx, "CRUD items")
	if res != true {
		return false, ErrForbidden
	}

	var updInput = &models.Item{
		ID:       input.ID,
		TypeID:   *input.TypeID,
		Lon:      *input.Lon,
		Lat:      *input.Lat,
		BrokenID: input.BrokenID,
	}

	return s.ItemsRepo.UpdateItem(updInput)
}

func (s *Services) DeleteItem(ctx context.Context, id string) (bool, error) {
	res := s.UsersRepo.AllowAccess(ctx, "CRUD items")
	if res != true {
		return false, ErrForbidden
	}

	itemEx := s.ItemsRepo.ItemInRentsExists(id)
	if itemEx == true {
		return false, ErrRelationExists
	}

	return s.ItemsRepo.DeleteItem(id)
}

func (s *Services) LocationGen(ctx context.Context, input models.LocationGenInput) (bool, error) {
	res := s.UsersRepo.AllowAccess(ctx, "CRUD items")
	if res != true {
		return false, ErrForbidden
	}

	if input.LonMin >= input.LonMax || input.LatMin >= input.LatMax {
		return false, errors.New("min value bigger than max")
	}

	loc, err := s.ItemsRepo.GetItems()
	if err != nil {
		return false, err
	}

	for _, locs := range loc {
		locs.Lon = input.LonMin + rand.Float64()*(input.LonMax-input.LonMin)
		locs.Lat = input.LatMin + rand.Float64()*(input.LatMax-input.LatMin)
	}

	return s.ItemsRepo.LocationGen(loc)
}

func (s *Services) GetRents(ctx context.Context) ([]*models.Rent, error) {
	res := s.UsersRepo.AllowAccess(ctx, "CRUD items")
	if res != true {
		return nil, ErrForbidden
	}

	return s.ItemsRepo.GetRents()
}

func (s *Services) GetRentsForItem(ctx context.Context, obj *models.Item) ([]*models.Rent, error) {
	res := s.UsersRepo.AllowAccess(ctx, "CRUD items")
	if res != true {
		return nil, ErrForbidden
	}

	return s.ItemsRepo.GetRentsForItem(obj)
}

func (s *Services) CreateRent(ctx context.Context, input models.CreateRentInput) (bool, error) {
	res := s.UsersRepo.AllowAccess(ctx, "CRUD items")
	if res != true {
		return false, ErrForbidden
	}

	itemEx := s.ItemsRepo.ItemExists(input.ItemsID)
	if itemEx == false {
		return false, ErrDatabaseValueNotExist
	}

	itemExInRents := s.ItemsRepo.ItemInRentsExists(input.ItemsID)
	if itemExInRents == true {
		return false, ErrDatabaseValueExists
	}

	currentUser := middleware.ForContext(ctx)

	rent := &models.Rent{
		ExternalID: input.ExternalID,
		ItemsID:    input.ItemsID,
		UsersID:    currentUser.ID,
	}

	return s.ItemsRepo.CreateRent(rent)
}

func (s *Services) DeleteRent(ctx context.Context, id string) (bool, error) {
	res := s.UsersRepo.AllowAccess(ctx, "CRUD items")
	if res != true {
		return false, ErrForbidden
	}

	return s.ItemsRepo.DeleteRent(id)
}
