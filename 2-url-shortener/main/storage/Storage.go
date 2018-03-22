package storage

type Storage interface {
	Code() string
	Save(string) string
	Load(string) (string, error)
}
