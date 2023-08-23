package main

import (
	"fmt"
	"http-method/handlers"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()

	router.GET("/inventories", handlers.GetInventories)
	router.GET("/inventories/:id", handlers.GetInventory)
	router.POST("/inventories", handlers.CreateInventory)
	router.PUT("/inventories/:id", handlers.UpdateInventory)
	router.DELETE("/inventories/:id", handlers.DeleteInventory)

	fmt.Println("Server is listening on 127.0.0.1:8080")
	http.ListenAndServe(":8080", router)
}
