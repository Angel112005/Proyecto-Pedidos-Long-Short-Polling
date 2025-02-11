package applicationpedido

import domainpedido "lab-test.com/module/src/pedidos/domain_pedido"

type DeletePedidoUseCase struct {
	repository domainpedido.PedidoRepository
}

func NewDeletePedidoUseCase(repository domainpedido.PedidoRepository) *DeletePedidoUseCase {
	return &DeletePedidoUseCase{repository: repository}
}

func (useCase *DeletePedidoUseCase) Execute(id int) error {
	return useCase.repository.Delete(id)
}