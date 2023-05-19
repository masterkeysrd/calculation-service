package record

import "errors"

type Service interface {
	FindByUserID(userID uint64) ([]Record, error)
	FindByUserIDAndID(userID uint64, id uint64) (Record, error)
	Create(record Record) error
	Delete(record Record) error
}

type recordService struct{}

func NewRecordService() Service {
	return &recordService{}
}

func (s *recordService) FindByUserID(userID uint64) ([]Record, error) {
	return []Record{}, errors.New("not implemented")
}

func (s *recordService) FindByUserIDAndID(userID uint64, id uint64) (Record, error) {
	return Record{}, errors.New("not implemented")
}

func (s *recordService) Create(record Record) error {
	return errors.New("not implemented")
}

func (s *recordService) Delete(record Record) error {
	return errors.New("not implemented")
}
