package main

import (
	"net/http"

	dimodule "github.com/alphadecodex/nomise-ragnarok/dimodule"
	"github.com/alphadecodex/nomise-ragnarok/resources"
	"github.com/gin-gonic/gin"
)

func main() {
	diModule := &dimodule.DIModule{}
	container := diModule.BuildContainer()

	err := container.Invoke(func(orderResource *resources.OrderResource) {
		router := gin.Default()
		router.POST("/orders", func(ctx *gin.Context) {
			var orderRequest resources.OrderRequest
			if err := ctx.ShouldBindJSON(&orderRequest); err != nil {
				// Handle the error
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			if err := orderResource.CreateOrder(ctx); err != nil {
				// Handle the error
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}

			ctx.JSON(http.StatusOK, gin.H{"message": "Order created successfully"})
		})

		// Start the server
		router.Run(":8080")
	})

	if err != nil {
		// Handle the error
	}
}
