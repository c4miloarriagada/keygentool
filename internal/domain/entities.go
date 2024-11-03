package domain

type Entities struct {
	Name   string
	Active bool
	Url    Url
	key    Keygen
}

type Url struct {
	Name   string
	HasUrl bool
}
