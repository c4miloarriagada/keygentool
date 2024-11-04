package domain

type EntityRepository interface {
	Save(entity Entities) error
	FindByID(id int) (Entities, error)
	FindAll() ([]Entities, error)
	Delete(id int) error
}
