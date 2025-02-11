package applicationpedido

import domainpedido "lab-test.com/module/src/pedidos/domain_pedido"

type ViewPedidoByIdUseCase struct {
	repository domainpedido.PedidoRepository
}

func NewViewPedidoByIdUseCase(repository domainpedido.PedidoRepository) *ViewPedidoByIdUseCase {
	return &ViewPedidoByIdUseCase{repository: repository}
}

func (useCase *ViewPedidoByIdUseCase) Execute(id int) (*domainpedido.Pedido, error) {
	return useCase.repository.ViewById(id)
}