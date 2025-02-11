package domainpedido

type PedidoRepository interface {
	Create(pedido *Pedido) error
	ViewAll() ([]*Pedido, error)
	Delete(id int) error
	Update(pedido *Pedido) error
	ViewById(id int) (*Pedido, error)
	ContarPendientes() (int, error) 
}