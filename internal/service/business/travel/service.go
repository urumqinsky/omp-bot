package travel

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) List() []Subdomain {
	return allEntities
}

func (s *Service) Get(idx int) (*Subdomain, error) {
	return &allEntities[idx], nil
}

func (s *Service) Add(title string) {
	allEntities = append(allEntities, NewSubdomain(title))
}

func (s *Service) Remove(idx int) {
	allEntities = append(allEntities[:idx], allEntities[idx+1:]...)
}
