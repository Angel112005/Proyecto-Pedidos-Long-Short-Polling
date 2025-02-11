package infrastructurepedido

import (
	"net/http"
	"strconv"
	"log"

	"github.com/gin-gonic/gin"
	applicationpedido "lab-test.com/module/src/pedidos/application_pedido"
	domainpedido "lab-test.com/module/src/pedidos/domain_pedido"
)

type PedidoController struct {
	createPedido    *applicationpedido.CreatePedidoUseCase
	viewPedidos     *applicationpedido.ViewAllPedidosUseCase
	viewPedidoById  *applicationpedido.ViewPedidoByIdUseCase
	updatePedido    *applicationpedido.UpdatePedidoUseCase
	deletePedido    *applicationpedido.DeletePedidoUseCase
	contarPendientes *applicationpedido.CountPedidosPendientesUseCase
}

func NewPedidoController(
	createPedido *applicationpedido.CreatePedidoUseCase, 
	viewPedidos *applicationpedido.ViewAllPedidosUseCase, 
	viewPedidoById *applicationpedido.ViewPedidoByIdUseCase, 
	updatePedido *applicationpedido.UpdatePedidoUseCase, 
	deletePedido *applicationpedido.DeletePedidoUseCase,
	contarPendientes *applicationpedido.CountPedidosPendientesUseCase) *PedidoController {
	
	return &PedidoController{
		createPedido:    createPedido,
		viewPedidos:     viewPedidos,
		viewPedidoById:  viewPedidoById,
		updatePedido:    updatePedido,
		deletePedido:    deletePedido,
		contarPendientes: contarPendientes,
	}
}

func (pc *PedidoController) CreatePedido(c *gin.Context) {
	var pedido domainpedido.Pedido
	if err := c.ShouldBindJSON(&pedido); err != nil {
		log.Println("❌ Error en la solicitud JSON:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := pc.createPedido.Execute(pedido.ClienteID, pedido.Estado); err != nil {
		log.Println("❌ Error creando pedido:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Println("✅ Pedido creado correctamente:", pedido)
	c.JSON(http.StatusOK, gin.H{"message": "Pedido creado correctamente"})
}

func (pc *PedidoController) GetAllPedidos(c *gin.Context) {
	pedidos, err := pc.viewPedidos.Execute()
	if err != nil {
		log.Println("❌ Error obteniendo pedidos:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Println("📌 Pedidos obtenidos con éxito")
	c.JSON(http.StatusOK, gin.H{"pedidos": pedidos})
}

// ✅ Obtener un pedido por ID
func (pc *PedidoController) GetPedidoById(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		log.Println("❌ Error: ID de pedido inválido", idParam)
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	pedido, err := pc.viewPedidoById.Execute(id)
	if err != nil {
		log.Println("❌ Error obteniendo pedido:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Println("📌 Pedido obtenido con éxito:", pedido)
	c.JSON(http.StatusOK, gin.H{"pedido": pedido})
}

// ✅ Actualizar un pedido
func (pc *PedidoController) UpdatePedido(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		log.Println("❌ Error: ID de pedido inválido", idParam)
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var pedido domainpedido.Pedido
	if err := c.ShouldBindJSON(&pedido); err != nil {
		log.Println("❌ Error en la solicitud JSON:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	pedido.ID = id
	if err := pc.updatePedido.Execute(&pedido); err != nil {
		log.Println("❌ Error actualizando pedido:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Println("✅ Pedido actualizado correctamente:", pedido)
	c.JSON(http.StatusOK, gin.H{"message": "Pedido actualizado correctamente"})
}

// ✅ Eliminar un pedido
func (pc *PedidoController) DeletePedido(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		log.Println("❌ Error: ID de pedido inválido", idParam)
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	if err := pc.deletePedido.Execute(id); err != nil {
		log.Println("❌ Error eliminando pedido:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Println("✅ Pedido eliminado con éxito, ID:", id)
	c.JSON(http.StatusOK, gin.H{"message": "Pedido eliminado correctamente"})
}

// ✅ Endpoint para contar pedidos pendientes (Short Polling)
func (pc *PedidoController) ContarPedidosPendientes(c *gin.Context) {
    cantidad, err := pc.contarPendientes.Execute()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"pedidos_pendientes": cantidad})
}