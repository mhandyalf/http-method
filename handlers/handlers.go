package handlers

import (
	"encoding/json"
	"http-method/config"
	"http-method/entity"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func GetInventories(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	inventories := config.FetchInventoriesFromDB()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(inventories)
}

func GetInventory(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	inventory := config.FetchInventoryFromDB(id)
	if inventory == nil {
		http.NotFound(w, r)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(inventory)
}

func CreateInventory(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var inventory entity.Inventory
	err := json.NewDecoder(r.Body).Decode(&inventory)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	config.InsertInventoryToDB(inventory)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(inventory)
}

func UpdateInventory(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	inventory := config.FetchInventoryFromDB(id)
	if inventory == nil {
		http.NotFound(w, r)
		return
	}

	var updatedInventory entity.Inventory
	err := json.NewDecoder(r.Body).Decode(&updatedInventory)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	updatedInventory.ID = inventory.ID
	config.UpdateInventoryInDB(updatedInventory)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedInventory)
}

func DeleteInventory(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	inventory := config.FetchInventoryFromDB(id)
	if inventory == nil {
		http.NotFound(w, r)
		return
	}

	config.DeleteInventoryFromDB(id)

	w.WriteHeader(http.StatusNoContent)
}
