package record

type RecordService struct{}

func NewRecordService() Service {
	return &RecordService{}
}

func (s *RecordService) FindByUserID(userID uint64) ([]Record, error) {
	return []Record{}, nil
}

func (s *RecordService) Create(record Record) error {
	return nil
}

func (s *RecordService) Delete(record Record) error {
	return nil
}
