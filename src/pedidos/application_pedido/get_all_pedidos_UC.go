package applicationpedido

import domainpedido "lab-test.com/module/src/pedidos/domain_pedido"

type ViewAllPedidosUseCase struct {
	repository domainpedido.PedidoRepository
}

func NewViewAllPedidosUseCase(repository domainpedido.PedidoRepository) *ViewAllPedidosUseCase {
	return &ViewAllPedidosUseCase{repository: repository}
}

func (useCase *ViewAllPedidosUseCase) Execute() ([]*domainpedido.Pedido, error) {
	return useCase.repository.ViewAll()
}