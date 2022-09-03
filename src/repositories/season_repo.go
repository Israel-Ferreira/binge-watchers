package repositories

type SeasonRepository interface {
	FindById(uint64) (interface{}, error)
	FindAll() (interface{}, error)
	DeleteById(uint64) error
	Create(interface{}) (interface{}, error)
	Update(uint64, interface{}) error
}
