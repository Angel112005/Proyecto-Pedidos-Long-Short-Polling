package domainpedido


type Pedido struct {
	ID          int    `json:"id"`
	ClienteID   int    `json:"cliente_id"`
	Estado      string `json:"estado"`
	RepartidorID int   `json:"repartidor_id"`
}

func NewPedido(clienteID int, estado string) *Pedido {
	return &Pedido{
		ClienteID: clienteID,
		Estado:    estado,
	}
}