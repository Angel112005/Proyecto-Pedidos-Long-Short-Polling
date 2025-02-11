package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"time"
	infrastructurepedido "lab-test.com/module/src/pedidos/infrastructure_pedido"
	infrastructurerepartidor "lab-test.com/module/src/repartidores/infrastructure_repartidor"
)

func main() {
	r := gin.Default()

	// Middleware de CORS para permitir solicitudes desde el frontend
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://127.0.0.1:5501"}, // Cambiar esto en caso de que el frontend tiene otra URL
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))


	// Inicialización de módulos de pedidos y repartidores
	infrastructurepedido.Init(r)
	infrastructurerepartidor.Init(r)

	// Ejecución del servidor en el puerto 3000
	if err := r.Run(":3000"); err != nil {
		panic(err)
	}
}
