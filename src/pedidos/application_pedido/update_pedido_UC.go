package applicationpedido

import domainpedido "lab-test.com/module/src/pedidos/domain_pedido"

type UpdatePedidoUseCase struct {
	repository domainpedido.PedidoRepository
}

func NewUpdatePedidoUseCase(repository domainpedido.PedidoRepository) *UpdatePedidoUseCase {
	return &UpdatePedidoUseCase{repository: repository}
}

func (useCase *UpdatePedidoUseCase) Execute(pedido *domainpedido.Pedido) error {
	return useCase.repository.Update(pedido)
}