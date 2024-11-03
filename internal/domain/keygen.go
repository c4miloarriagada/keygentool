package domain

import "time"

type Keygen struct {
	HashedPwd string
	Active    bool
	CreatedAt time.Time
}
