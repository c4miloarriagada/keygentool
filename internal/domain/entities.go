package domain

type Entities struct {
	Name   string
	Active bool
	Url    Url
}

type Url struct {
	Name   string
	HasUrl bool
}
