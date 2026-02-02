package ports

type Irepository[T any] interface {
	Save(info T) (int64, error)
	Get(id int64) (T, error)
	GetList() ([]T, error)
	Edit(info T) error
	Delete(id int64) error
}
