package travel

var allEntities []Subdomain

type Subdomain struct {
	Title string
}

func NewSubdomain(name string) Subdomain {
	return Subdomain{Title: name}
}
