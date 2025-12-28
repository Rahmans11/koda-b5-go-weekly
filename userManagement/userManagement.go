package usermanagement

import "fmt"

type User struct {
	Id       string
	Username string
}

type UserManager struct {
	Users map[string]*User
}

func NewUserManager() *UserManager {
	return &UserManager{
		Users: make(map[string]*User),
	}
}

func (um *UserManager) AddUser(user User) *UserManager {
	if _, id := um.Users[user.Id]; id {
		fmt.Println("Failed: Id sudah digunakan")
		return um
	}
	um.Users[user.Id] = &user
	return um
}

func (um *UserManager) GetUserById(id string) (*User, bool) {
	user, exists := um.Users[id]
	return user, exists
}
