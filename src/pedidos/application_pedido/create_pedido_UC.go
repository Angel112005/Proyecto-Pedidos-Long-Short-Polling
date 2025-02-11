package applicationpedido

import domainpedido "lab-test.com/module/src/pedidos/domain_pedido"

type CreatePedidoUseCase struct {
	repository domainpedido.PedidoRepository
}

func NewCreatePedidoUseCase(repository domainpedido.PedidoRepository) *CreatePedidoUseCase {
	return &CreatePedidoUseCase{repository: repository}
}

func (useCase *CreatePedidoUseCase) Execute(clienteID int, estado string) error {
	pedido := domainpedido.NewPedido(clienteID, estado)
	return useCase.repository.Create(pedido)
}