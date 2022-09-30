package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"text/template"
)

type Employee struct {
	ID int
	Name string
	Age int
	Division string
}

var employees = []Employee{
	{ID: 1, Name: "renny", Age: 20, Division:  "IT"},
	{ID: 2, Name: "kristina", Age:  21, Division:  "finance"},
	{ID: 3, Name: "reny", Age:  23, Division: "Akun"},
}

var PORT = ":8080"

func main()  {
	http.HandleFunc("/employees", getEmployees)

	http.HandleFunc("/employee", createEmployee)

	fmt.Println("Aplication is listening on port", PORT)
	http.ListenAndServe(PORT, nil)
}

func getEmployees(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
	tpl, err := template.ParseFiles("template.html")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tpl.Execute(w, employees)
	return
	}

	http.Error(w, "invalid method", http.StatusBadRequest)
}

func createEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "aplication/json")

	if r.Method == "POST" {
		name := r.FormValue("name")
		age := r.FormValue("age")
		division := r.FormValue("division")

		convertAge, err := strconv.Atoi(age)

		if err != nil {
			http.Error(w, "invalid age", http.StatusBadRequest)
			return
		}
		newEmployee := Employee {
		ID: len(employees) + 1,
		Name: name,
		Age: convertAge,
		Division: division,
	}	

	employees = append(employees, newEmployee)
	
	json.NewEncoder(w).Encode(newEmployee)
	return
	}

	http.Error(w, "invalid method", http.StatusBadRequest)
}
