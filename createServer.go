package main

import (
	router "createServer/routers"
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World!")
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/add", router.AdditionHandler)
	http.HandleFunc("/items", router.GetItemHandler)
	http.HandleFunc("/items/add", router.AddItemHandler)
	http.HandleFunc("/items/update", router.UpdateItemHandler)
	http.HandleFunc("/items/delete", router.DeleteItemHandler)

	fmt.Println("Server listening on port: 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
