package service

type Hasher interface {
	Hash(baseString string) (string, error)
	NewUID(userLogin string) (string, error)
}

type Service struct {
	Hasher
}

func New() *Service {
	return &Service{
		Hasher: NewHasherService(),
	}
}
