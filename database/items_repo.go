package database

import (
	"errors"
	"github.com/go-pg/pg"
	"trackingApp/graph/models"
)

type ItemsRepo struct {
	DB *pg.DB
}

func (i *ItemsRepo) CreateItem(item *models.Item) (bool, error) {
	_, err := i.DB.Model(item).Insert()
	if err != nil {
		return false, errors.New("can't create Item")
	}

	return true, nil
}

func (i *ItemsRepo) GetItemByID(id string) (*models.Item, error) {
	var item models.Item
	err := i.DB.Model(&item).Where("id = ?", id).First()
	if err != nil {
		return nil, err
	}

	return &item, nil
}

func (i *ItemsRepo) GetItems() ([]*models.Item, error) {
	var items []*models.Item
	err := i.DB.Model(&items).Where("delete_status = 1").Order("id").Select()
	if err != nil {
		return nil, errors.New("can't get Items")
	}

	return items, nil
}

func (i *ItemsRepo) GetInputTypes() ([]*models.InputTypes, error) {
	var inputTypes []*models.InputTypes
	_, err := i.DB.Query(&inputTypes, `
	SELECT i.id, it.name, i.title
	FROM inputs i
	JOIN input_types it ON it.id = i.input_types_id
	`)

	if err != nil {
		return nil, errors.New("can't get input types")
	}

	return inputTypes, nil
}

func (i *ItemsRepo) GetTypeNames() ([]*models.TypeNames, error) {
	var typeNames []*models.TypeNames
	_, err := i.DB.Query(&typeNames, `
	SELECT name	FROM input_types
	`)

	if err != nil {
		return nil, errors.New("can't get type names")
	}

	return typeNames, nil
}

func (i *ItemsRepo) UpdateItem(input *models.Item) (bool, error) {
	res, err := i.DB.Model(input).WherePK().UpdateNotNull()
	if err != nil || res.RowsAffected() == 0 {
		return false, errors.New("can't update Item")
	}

	return true, nil
}

func (i *ItemsRepo) LocationGen(input []*models.Item) (bool, error) {
	for _, locs := range input {
		res, _ := i.DB.Model(locs).WherePK().UpdateNotNull()
		if res.RowsAffected() == 0 {
			return false, errors.New("can't update Location")
		}
	}

	return true, nil
}

func (i *ItemsRepo) DeleteItem(id string) (bool, error) {
	var item *models.Item
	_, err := i.GetItemByID(id)
	if err != nil {
		return false, errors.New("can't delete Item")
	}

	_, err = i.DB.Model(item).Set("delete_status = 0").Where("id = ?", id).Update()
	if err != nil {
		return false, errors.New("error deleting item")
	}

	return true, nil
}

func (i *ItemsRepo) GetRents() ([]*models.Rent, error) {
	var rents []*models.Rent
	err := i.DB.Model(&rents).Where("delete_status = 1").Order("id").Select()
	if err != nil {
		return nil, errors.New("can't get Rents")
	}

	return rents, nil
}

func (i *ItemsRepo) GetRentsForItem(obj *models.Item) ([]*models.Rent, error) {
	var rents []*models.Rent
	err := i.DB.Model(&rents).Where("items_id = ?", obj.ID).Where("delete_status = 1").Select()
	if err != nil {
		return nil, err
	}

	return rents, err
}

func (i *ItemsRepo) CreateRent(input *models.Rent) (bool, error) {
	_, err := i.DB.Model(input).Insert()
	if err != nil {
		return false, errors.New("can't create Rent")
	}

	return true, nil
}

func (i *ItemsRepo) DeleteRent(id string) (bool, error) {
	var rent *models.Rent
	_, err := i.DB.Model(rent).Set("delete_status = 0").Where("id = ?", id).Update()
	if err != nil {
		return false, errors.New("error deleting item")
	}

	return true, nil
}

func (i *ItemsRepo) ItemInRentsExists(id string) bool {
	var rent *models.Rent
	res, _ := i.DB.Model(rent).Where("items_id = ?", id).Where("delete_status = 1").Count()
	if res != 1 {
		return false
	}

	return true
}

func (i *ItemsRepo) ItemExists(id string) bool {
	var item *models.Item
	res, _ := i.DB.Model(item).Where("id = ?", id).Where("delete_status = 1").Count()
	if res != 1 {
		return false
	}

	return true
}
