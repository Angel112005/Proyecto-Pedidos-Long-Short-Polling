package infrastructurerepartidor

import (
	"github.com/gin-gonic/gin"
	applicationrepartidor "lab-test.com/module/src/repartidores/application_repartidor"
)

func Init(r *gin.Engine) {
	ps := NewMySQLRepartidor()
	createRepartidor := applicationrepartidor.NewCreateRepartidorUseCase(ps)
	viewRepartidores := applicationrepartidor.NewViewAllRepartidoresUseCase(ps)
	viewRepartidorById := applicationrepartidor.NewViewRepartidorByIdUseCase(ps)
	updateRepartidor := applicationrepartidor.NewUpdateRepartidorUseCase(ps)
	deleteRepartidor := applicationrepartidor.NewDeleteRepartidorUseCase(ps)

	repartidorController := NewRepartidorController(createRepartidor, viewRepartidores, viewRepartidorById, updateRepartidor, deleteRepartidor)
	RegisterRepartidorRoutes(r, repartidorController)
}
