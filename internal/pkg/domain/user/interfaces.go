package user

type Service interface {
	FindByUserName(userName string) (User, error)
	FindByID(id uint64) (User, error)
	Create(user User) error
	Delete(id uint64) error
}
