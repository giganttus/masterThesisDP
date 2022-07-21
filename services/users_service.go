package services

import (
	"context"
	"errors"
	"trackingApp/graph/models"
	"trackingApp/middleware"
)

func (s *Services) ProfileSettings(ctx context.Context) (*models.UserSettings, error) {
	currentUser := middleware.ForContext(ctx)
	if currentUser == nil {
		return nil, errors.New("Unauthenticated")
	}

	return s.UsersRepo.GetUserSettings(currentUser.ID)
}

func (s *Services) UserSettings(ctx context.Context, id string) (*models.UserSettings, error) {
	res := s.UsersRepo.AllowAccess(ctx, "User settings")
	if res != true {
		return nil, ErrForbidden
	}

	return s.UsersRepo.GetUserSettings(id)
}

func (s *Services) CreateUser(ctx context.Context, input models.CreateUserInput) (bool, error) {
	res := s.UsersRepo.AllowAccess(ctx, "CRUD users")
	if res != true {
		return false, ErrForbidden
	}

	userEx := s.UsersRepo.UsernameExists(input.Username)
	if userEx == true {
		return false, ErrDatabaseValueExists
	}

	if len(input.Username) > 50 || len(input.FirstName) > 50 || len(input.LastName) > 50 || len(input.Password) > 50 || len(input.Email) > 50 {
		return false, ErrLengthLimit
	}

	hashedPwd, err := models.HashPassword(input.Password)
	if err != nil {
		return false, ErrPasswordHash
	}
	input.Password = hashedPwd

	var creInput = &models.User{
		Username:  input.Username,
		Email:     input.Email,
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Password:  input.Password,
		RolesID:   input.RolesID,
	}

	return s.UsersRepo.CreateUser(creInput)
}

func (s *Services) GetUsers(ctx context.Context) ([]*models.User, error) {
	res := s.UsersRepo.AllowAccess(ctx, "CRUD users")
	if res != true {
		return nil, ErrForbidden
	}

	return s.UsersRepo.GetUsers()
}

func (s *Services) UpdateUser(ctx context.Context, input models.UpdateUserInput) (bool, error) {
	res := s.UsersRepo.AllowAccess(ctx, "Personal update")
	if res != true {
		return false, ErrForbidden
	}

	var updInput = &models.User{
		ID:        input.ID,
		FirstName: *input.FirstName,
		LastName:  *input.LastName,
	}

	return s.UsersRepo.UpdateUser(updInput)
}

func (s *Services) UpdateAdmin(ctx context.Context, input models.UpdateAdminInput) (bool, error) {
	res := s.UsersRepo.AllowAccess(ctx, "CRUD users")
	if res != true {
		return false, ErrForbidden
	}

	var updInput = &models.User{
		ID:        input.ID,
		FirstName: *input.FirstName,
		LastName:  *input.LastName,
		RolesID:   input.RolesID,
	}

	return s.UsersRepo.UpdateAdmin(updInput)
}

func (s *Services) DeleteUser(ctx context.Context, id string) (bool, error) {
	res := s.UsersRepo.AllowAccess(ctx, "CRUD users")
	if res != true {
		return false, ErrForbidden
	}

	return s.UsersRepo.DeleteUser(id)
}

func (s *Services) ActivateUser(ctx context.Context, id string) (bool, error) {
	res := s.UsersRepo.AllowAccess(ctx, "CRUD users")
	if res != true {
		return false, ErrForbidden
	}

	return s.UsersRepo.ActivateUser(id)
}

func (s *Services) ChangePassword(ctx context.Context, input models.ChangePasswordInput) (bool, error) {
	res := s.UsersRepo.AllowAccess(ctx, "CRUD users")
	if res != true {
		return false, ErrForbidden
	}

	hashedPwd, err := models.HashPassword(input.Password)
	if err != nil {
		return false, ErrPasswordHash
	}
	input.Password = hashedPwd

	var cpwInput = &models.User{
		ID:       input.ID,
		Password: input.Password,
	}

	return s.UsersRepo.ChangePassword(cpwInput)
}

func (s *Services) GetRoleNames(ctx context.Context) ([]*models.Role, error) {

	res := s.UsersRepo.AllowAccess(ctx, "CRUD users")
	if res != true {
		return nil, ErrForbidden
	}
	return s.UsersRepo.GetRoleNames()
}
