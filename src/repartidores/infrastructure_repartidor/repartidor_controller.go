package infrastructurerepartidor

import (
	"net/http"
	"strconv"
	"log"

	"github.com/gin-gonic/gin"
	applicationrepartidor "lab-test.com/module/src/repartidores/application_repartidor"
	domainrepartidor "lab-test.com/module/src/repartidores/domain_repartidor"
)

type RepartidorController struct {
	createRepartidor   *applicationrepartidor.CreateRepartidorUseCase
	viewRepartidores   *applicationrepartidor.ViewAllRepartidoresUseCase
	viewRepartidorById *applicationrepartidor.ViewRepartidorByIdUseCase
	updateRepartidor   *applicationrepartidor.UpdateRepartidorUseCase
	deleteRepartidor   *applicationrepartidor.DeleteRepartidorUseCase
}

func NewRepartidorController(
	createRepartidor *applicationrepartidor.CreateRepartidorUseCase,
	viewRepartidores *applicationrepartidor.ViewAllRepartidoresUseCase,
	viewRepartidorById *applicationrepartidor.ViewRepartidorByIdUseCase,
	updateRepartidor *applicationrepartidor.UpdateRepartidorUseCase,
	deleteRepartidor *applicationrepartidor.DeleteRepartidorUseCase) *RepartidorController {

	return &RepartidorController{
		createRepartidor:   createRepartidor,
		viewRepartidores:   viewRepartidores,
		viewRepartidorById: viewRepartidorById,
		updateRepartidor:   updateRepartidor,
		deleteRepartidor:   deleteRepartidor,
	}
}

// ✅ Crear un nuevo repartidor
func (rc *RepartidorController) CreateRepartidor(c *gin.Context) {
	var repartidor domainrepartidor.Repartidor
	if err := c.ShouldBindJSON(&repartidor); err != nil {
		log.Println("❌ Error en la solicitud JSON:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := rc.createRepartidor.Execute(repartidor.Nombre, repartidor.Disponible); err != nil {
		log.Println("❌ Error creando repartidor:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Println("✅ Repartidor creado correctamente:", repartidor)
	c.JSON(http.StatusOK, gin.H{"message": "Repartidor creado correctamente"})
}

// ✅ Obtener todos los repartidores
func (rc *RepartidorController) GetAllRepartidores(c *gin.Context) {
	repartidores, err := rc.viewRepartidores.Execute()
	if err != nil {
		log.Println("❌ Error obteniendo repartidores:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Println("📌 Repartidores obtenidos con éxito")
	c.JSON(http.StatusOK, gin.H{"repartidores": repartidores})
}

// ✅ Obtener un repartidor por ID
func (rc *RepartidorController) GetRepartidorById(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		log.Println("❌ Error: ID de repartidor inválido", idParam)
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	repartidor, err := rc.viewRepartidorById.Execute(id)
	if err != nil {
		log.Println("❌ Error obteniendo repartidor:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Println("📌 Repartidor obtenido con éxito:", repartidor)
	c.JSON(http.StatusOK, gin.H{"repartidor": repartidor})
}

// ✅ Actualizar un repartidor
func (rc *RepartidorController) UpdateRepartidor(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		log.Println("❌ Error: ID de repartidor inválido", idParam)
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var repartidor domainrepartidor.Repartidor
	if err := c.ShouldBindJSON(&repartidor); err != nil {
		log.Println("❌ Error en la solicitud JSON:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	repartidor.ID = id
	if err := rc.updateRepartidor.Execute(&repartidor); err != nil {
		log.Println("❌ Error actualizando repartidor:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Println("✅ Repartidor actualizado correctamente:", repartidor)
	c.JSON(http.StatusOK, gin.H{"message": "Repartidor actualizado correctamente"})
}

// ✅ Eliminar un repartidor
func (rc *RepartidorController) DeleteRepartidor(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		log.Println("❌ Error: ID de repartidor inválido", idParam)
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	if err := rc.deleteRepartidor.Execute(id); err != nil {
		log.Println("❌ Error eliminando repartidor:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Println("✅ Repartidor eliminado con éxito, ID:", id)
	c.JSON(http.StatusOK, gin.H{"message": "Repartidor eliminado correctamente"})
}
