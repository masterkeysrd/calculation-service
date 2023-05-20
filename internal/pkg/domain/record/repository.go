package record

import "time"

var records = []Record{
	{
		ID:          1,
		UserID:      1,
		OperationID: 1,
		Amount:      100,
		UserBalance: 100,
		CreatedAt:   time.Now(),
	},
}

type Repository interface {
	FindByUserID(userID uint64) (*[]Record, error)
	FindByUserIDAndID(userID uint64, id uint64) (*Record, error)
	Create(record *Record) error
	Update(record *Record) error
	Delete(record *Record) error
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) FindByUserID(userID uint64) (*[]Record, error) {
	result := []Record{}
	for _, record := range records {
		if record.UserID == userID {
			result = append(result, record)
		}
	}
	return &result, nil
}

func (r *repository) FindByUserIDAndID(userID uint64, id uint64) (*Record, error) {
	for _, record := range records {
		if record.UserID == userID && record.ID == id {
			return &record, nil
		}
	}
	return nil, nil
}

func (r *repository) Create(record *Record) error {
	records = append(records, *record)
	return nil
}

func (r *repository) Update(record *Record) error {
	for i, r := range records {
		if r.UserID == record.UserID && r.ID == record.ID {
			records[i] = *record
			return nil
		}
	}
	return nil
}

func (r *repository) Delete(record *Record) error {
	for i, r := range records {
		if r.UserID == record.UserID && r.ID == record.ID {
			records = append(records[:i], records[i+1:]...)
			return nil
		}
	}
	return nil
}
