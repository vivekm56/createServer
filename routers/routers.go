package router

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type Item struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

var items []Item

func AdditionHandler(w http.ResponseWriter, r *http.Request) {
	a, errA := strconv.Atoi(r.URL.Query().Get("a"))
	b, errB := strconv.Atoi(r.URL.Query().Get("b"))

	if errA != nil || errB != nil {
		http.Error(w, "Invlid input", http.StatusBadRequest)
	}
	sum := a + b
	fmt.Fprintf(w, "Result %d\n", sum)
}

func AddItemHandler(w http.ResponseWriter, r *http.Request) {
	var newItem Item
	err := json.NewDecoder(r.Body).Decode(&newItem)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	newItem.Id = len(items) + 1
	items = append(items, newItem)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newItem)

}

func GetItemHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}

func UpdateItemHandler(w http.ResponseWriter, r *http.Request) {
	var updatedItem Item
	err := json.NewDecoder(r.Body).Decode(&updatedItem)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	for i, item := range items {
		if item.Id == updatedItem.Id {
			items[i] = updatedItem

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(updatedItem)
			return
		}
	}
	http.Error(w, "Item not found", http.StatusNotFound)
}

func DeleteItemHandler(w http.ResponseWriter, r *http.Request) {
	var itemToDelete Item
	err := json.NewDecoder(r.Body).Decode(&itemToDelete)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	for i, item := range items {
		if item.Id == itemToDelete.Id {
			items = append(items[:i], items[i+1:]...)
			w.WriteHeader(http.StatusOK)
			return
		}
	}
	http.Error(w, "Item not found", http.StatusNotFound)
}
