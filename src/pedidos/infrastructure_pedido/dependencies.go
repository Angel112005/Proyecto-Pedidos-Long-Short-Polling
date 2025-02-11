package infrastructurepedido

import (
	"github.com/gin-gonic/gin"
	applicationpedido "lab-test.com/module/src/pedidos/application_pedido"

)

func Init(r *gin.Engine) {
	ps := NewMySQLPedido()
	createPedido := applicationpedido.NewCreatePedidoUseCase(ps)
	viewPedidos := applicationpedido.NewViewAllPedidosUseCase(ps)
	viewPedidoById := applicationpedido.NewViewPedidoByIdUseCase(ps)
	updatePedido := applicationpedido.NewUpdatePedidoUseCase(ps)
	deletePedido := applicationpedido.NewDeletePedidoUseCase(ps)
	contarPendientes := applicationpedido.NewCountPedidosPendientesUseCase(ps) // ✅ Nuevo caso de uso


	pedidoController := NewPedidoController(createPedido, viewPedidos, viewPedidoById, updatePedido, deletePedido, contarPendientes)
	RegisterPedidoRoutes(r, pedidoController)
}
