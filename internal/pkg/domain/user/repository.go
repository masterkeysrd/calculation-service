package user

type Repository interface {
	FindByID(id uint) (*User, error)
	FindByUserName(userName string) (*User, error)
	Create(user *User) error
	Update(user *User) error
	Delete(user *User) error
}
