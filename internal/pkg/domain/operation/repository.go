package operation

type Repository interface {
	List() ([]*Operation, error)
	Get(id uint) (*Operation, error)
}
