package infrastructurepedido

import "github.com/gin-gonic/gin"

func RegisterPedidoRoutes(r *gin.Engine, pedidoController *PedidoController) {
	pedidos := r.Group("/pedido")
	{
		pedidos.POST("", pedidoController.CreatePedido)
		pedidos.GET("", pedidoController.GetAllPedidos)
		pedidos.GET("/:id", pedidoController.GetPedidoById)
		pedidos.GET("/pendientes", pedidoController.ContarPedidosPendientes)
		pedidos.PUT("/:id", pedidoController.UpdatePedido)
		pedidos.DELETE("/:id", pedidoController.DeletePedido)
	}
}