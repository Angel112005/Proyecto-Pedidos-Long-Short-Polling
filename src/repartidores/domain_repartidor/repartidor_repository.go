package domainrepartidor

type RepartidorRepository interface {
	Create(repartidor *Repartidor) error
	ViewAll() ([]*Repartidor, error)
	Delete(id int) error
	Update(repartidor *Repartidor) error
	ViewById(id int) (*Repartidor, error)
}
