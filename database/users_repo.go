package database

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-pg/pg"
	"github.com/julianfrank/console"
	"trackingApp/graph/models"
	"trackingApp/middleware"
)

type UsersRepo struct {
	DB *pg.DB
}

func (u *UsersRepo) GetUserByField(field, value string) (*models.User, error) {
	var user models.User
	err := u.DB.Model(&user).Where(fmt.Sprintf("%v = ?", field), value).First()
	return &user, err
}

func (u *UsersRepo) GetUserByID(id string) (*models.User, error) {
	return u.GetUserByField("id", id)
}

func (u *UsersRepo) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := u.DB.Model(&user).Where("email = ?", email).First()

	if err != nil {
		return nil, errors.New("can't get User")
	}

	return &user, nil
}

func (u *UsersRepo) GetUserByUsername(username string) (*models.User, error) {
	return u.GetUserByField("username", username)
}

func (u *UsersRepo) GetRole(userId string) (string, error) {
	var user *models.User
	var roleId string
	query := u.DB.Model(user).Column("roles_id").Where("id = ?", userId).Select(&roleId)

	if query != nil {
		return "", errors.New("no user with that id")

	}

	return roleId, nil
}

func (u *UsersRepo) GetUserSettings(id string) (*models.UserSettings, error) {
	var user *models.User
	var username, firstname, lastname, email string
	query := u.DB.Model(user).Column("username", "first_name", "last_name", "email").
		Where("id = ?", id).Select(&username, &firstname, &lastname, &email)

	if query != nil {
		return nil, errors.New("No ")
	}

	userSett := &models.UserSettings{
		Username:  username,
		Email:     email,
		FirstName: firstname,
		LastName:  lastname,
	}

	return userSett, nil
}

func (u *UsersRepo) CreateUser(input *models.User) (bool, error) {
	_, err := u.DB.Model(input).Insert()

	if err != nil {
		return false, errors.New("cant create User")
	}

	return true, nil
}

func (u *UsersRepo) GetUsers() ([]*models.User, error) {
	var users []*models.User
	err := u.DB.Model(&users).Order("id").Select()

	if err != nil {
		return nil, errors.New("can't get Users")
	}

	return users, nil
}

func (u *UsersRepo) UpdateUser(input *models.User) (bool, error) {
	res, err := u.DB.Model(input).WherePK().UpdateNotNull()

	if err != nil || res.RowsAffected() == 0 {
		return false, errors.New("can't update User")
	}

	return true, nil
}

func (u *UsersRepo) UpdateAdmin(input *models.User) (bool, error) {
	res, err := u.DB.Model(input).WherePK().UpdateNotNull()

	if err != nil || res.RowsAffected() == 0 {
		return false, errors.New("can't update User")
	}

	return true, nil
}

func (u *UsersRepo) DeleteUser(id string) (bool, error) {
	var user *models.User
	_, err := u.GetUserByID(id)

	if err != nil {
		return false, errors.New("can't delete user")
	}

	res, updErr := u.DB.Model(user).Set("delete_status = 0").
		Where("delete_status = 1").Where("id = ?", id).Update()

	if updErr != nil || res.RowsAffected() == 0 {
		return false, errors.New("can't delete User")
	}

	return true, nil
}

func (u *UsersRepo) ActivateUser(id string) (b bool, bool error) {
	var user *models.User
	_, err := u.GetUserByID(id)
	if err != nil {
		return false, errors.New("can't activate User")
	}
	_, err = u.DB.Model(user).Set("delete_status = 1").Where("id = ?", id).Update()

	return true, nil
}

func (u *UsersRepo) ChangePassword(input *models.User) (bool, error) {
	res, err := u.DB.Model(input).WherePK().UpdateNotNull()

	if err != nil || res.RowsAffected() == 0 {
		return false, errors.New("can't change password")
	}

	return true, nil
}

func (u *UsersRepo) AllowAccess(ctx context.Context, p string) bool {
	currentUser := middleware.ForContext(ctx)
	var permissions *models.Permissions
	var rolesPermissions *models.RolesPermissions
	var permId int

	if currentUser == nil {
		console.Log("Unauthenticated")
		return false
	}

	userRole, _ := u.GetRole(currentUser.ID)
	query := u.DB.Model(permissions).Column("id").Where("title = ?", p).Select(&permId)

	if query != nil {
		return false
	}

	res, _ := u.DB.Model(rolesPermissions).Where("roles_id = ?", userRole).
		Where("permissions_id = ?", permId).Exists()

	if res != true {
		console.Log("Permission not found for that user")
		return false
	}

	return true
}

func (u *UsersRepo) UsernameExists(username string) bool {
	var user *models.User
	res, _ := u.DB.Model(user).Where("username = ?", username).Count()

	if res != 1 {
		return false
	}

	return true
}

func (u *UsersRepo) GetRoleNames() ([]*models.Role, error) {
	var roles []*models.Role
	err := u.DB.Model(&roles).Select()

	if err != nil {
		return nil, errors.New("can't get Role names")
	}

	return roles, nil
}

func (u *UsersRepo) IsActive(id string) bool {
	var user *models.User
	res, _ := u.DB.Model(user).
		Where("id=?", id).
		Where("delete_status = ?", 0).Count()

	if res != 0 {
		return false
	}
	return true
}
