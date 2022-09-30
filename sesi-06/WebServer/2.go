package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Employee struct {
	ID int
	Name string
	Age string
	Division string
}

var employees = []Employee{
	{ID: 1, Name: "renny", Age:  20, Division:  "IT"},
	{ID: 2, Name: "kristina", Age:  21, Division:  "finance"},
	{ID: 3, Name: "reny", Age:  23, Division: "Akun"},
}

var PORT = ":8080"

func main()  {
	http.HandleFunc("/employees", getEmployees)

	fmt.Println("Aplication is listening on port", PORT)
	http.ListenAndServe(PORT, nil)
}

func getEmployees(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {
		json.NewEncoder (w).Encode(employees)
		return
	}
	http.Error(w, "invalid method", http.StatusBadRequest)
}