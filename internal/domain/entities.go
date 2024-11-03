package domain

type IDGenerator interface {
	NextID() int
}

type Entities struct {
	ID     int
	Name   string
	Active bool
	Url    Url
	Key    Keygen
}

type Url struct {
	Name   string
	HasUrl bool
}

func NewEntity(name string, active bool, url Url, key Keygen, idGen IDGenerator) Entities {
	return Entities{
		ID:     idGen.NextID(),
		Name:   name,
		Active: active,
		Url:    url,
		Key:    key,
	}
}
