package domain

import "time"

type Keygen struct {
	HashedPwd string
	Active    bool
	CreatedAt time.Time
}

func (e *Keygen) CreateNewKeygen() {
	e.Active = false
	e.HashedPwd = ""
	e.CreatedAt = time.Now()
}
