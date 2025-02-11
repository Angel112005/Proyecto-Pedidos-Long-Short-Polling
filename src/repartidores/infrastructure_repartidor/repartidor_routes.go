package infrastructurerepartidor

import "github.com/gin-gonic/gin"

func RegisterRepartidorRoutes(r *gin.Engine, repartidorController *RepartidorController) {
	repartidores := r.Group("/repartidores")
	{
		repartidores.POST("", repartidorController.CreateRepartidor)
		repartidores.GET("", repartidorController.GetAllRepartidores)
		repartidores.GET("/:id", repartidorController.GetRepartidorById)
		repartidores.PUT("/:id", repartidorController.UpdateRepartidor)
		repartidores.DELETE("/:id", repartidorController.DeleteRepartidor)
	}
}
