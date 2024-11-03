package utils

import "sync"

type autoInc struct {
	sync.Mutex
	id int
}

func (a *autoInc) NextID() (id int) {
	a.Lock()
	defer a.Unlock()

	id = a.id
	a.id++
	return
}
