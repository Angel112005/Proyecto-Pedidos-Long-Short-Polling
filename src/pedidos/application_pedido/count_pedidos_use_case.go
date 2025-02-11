package applicationpedido

import domainpedido "lab-test.com/module/src/pedidos/domain_pedido"

type CountPedidosPendientesUseCase struct {
	repository domainpedido.PedidoRepository
}

func NewCountPedidosPendientesUseCase(repository domainpedido.PedidoRepository) *CountPedidosPendientesUseCase {
	return &CountPedidosPendientesUseCase{repository: repository}
}

func (uc *CountPedidosPendientesUseCase) Execute() (int, error) {
	return uc.repository.ContarPendientes()
}
