package user

import (
	"errors"
	"todoapp/internal/core/helpers"
	entities "todoapp/internal/entities"
)

type UserRepository struct {
	users []*entities.User
}

func init() {

}
func (u *UserRepository) Add(user entities.User) {
	user.ID = helpers.LastIndex(u.users, "ID").(int) + 1

	u.users = append(u.users, &user)

}
func Delete(usr entities.User) {}
func (u *UserRepository) FindByUser(usr entities.User) (*entities.User, error) {
	for i, user := range u.users {
		if user.UserName == usr.UserName && user.Password == usr.Password {
			return u.users[i], nil
		}
	}
	return nil, errors.New("couldnt find")
}

func NewUserRepository(users []*entities.User) *UserRepository {
	return &UserRepository{users: users}
}
