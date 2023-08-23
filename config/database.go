package config

import (
	"database/sql"
	"http-method/entity"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/avenger")
	if err != nil {
		panic(err)
	}
}

func FetchInventoriesFromDB() []entity.Inventory {
	rows, err := db.Query("SELECT * FROM inventory")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	inventories := []entity.Inventory{}
	for rows.Next() {
		var inventory entity.Inventory
		err := rows.Scan(&inventory.ID, &inventory.Name, &inventory.Item_code, &inventory.Stock, &inventory.Description, &inventory.Status)
		if err != nil {
			panic(err)
		}
		inventories = append(inventories, inventory)
	}

	return inventories
}

func FetchInventoryFromDB(id string) *entity.Inventory {
	var inventory entity.Inventory
	row := db.QueryRow("SELECT * FROM inventory WHERE id = ?", id)
	err := row.Scan(&inventory.ID, &inventory.Name, &inventory.Item_code, &inventory.Stock, &inventory.Description, &inventory.Status)
	if err == sql.ErrNoRows {
		return nil
	} else if err != nil {
		panic(err)
	}

	return &inventory
}

func InsertInventoryToDB(inventory entity.Inventory) {
	_, err := db.Exec("INSERT INTO inventory (name, item_code, stock, description, status) VALUES (?, ?, ?, ?, ?)", inventory.Name, inventory.Item_code, inventory.Stock, inventory.Description, inventory.Status)
	if err != nil {
		panic(err)
	}
}

func UpdateInventoryInDB(inventory entity.Inventory) {
	_, err := db.Exec("UPDATE inventory SET name = ?, description = ?, stock = ? WHERE id = ?", inventory.Name, inventory.Description, inventory.Stock, inventory.ID)
	if err != nil {
		panic(err)
	}
}

func DeleteInventoryFromDB(id string) {
	_, err := db.Exec("DELETE FROM inventory WHERE id = ?", id)
	if err != nil {
		panic(err)
	}
}
