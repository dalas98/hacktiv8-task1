package main

import (
	model "dalas98/golang-lesson/Task-7/model"

	"encoding/json"
	"fmt"
	"net/http"
)

var port = ":8899"

func main() {
	http.HandleFunc("/", greeting)
	http.HandleFunc("/get-employee", getEmployees)
	http.HandleFunc("/create-employee", createEmployee)

	fmt.Println("Listening on port", port)
	http.ListenAndServe(port, nil)
}

func greeting(w http.ResponseWriter, r *http.Request) {
	message := "Hello World"
	fmt.Fprint(w, message)
}

func getEmployees(w http.ResponseWriter, r *http.Request) {
	employee := model.GetEmployee()
	json.NewEncoder(w).Encode(map[string]interface{}{
		"payload": employee,
		"total":   len(employee),
	})
}

func createEmployee(w http.ResponseWriter, r *http.Request) {
	method := r.Method

	if method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)

		json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "method " + method + " not allowed",
		})

		return
	}

	var req model.Employee
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		json.NewEncoder(w).Encode(map[string]interface{}{
			"error": err.Error(),
		})

		return
	}

	req.ID = len(model.Emps) + 1
	create := model.CreateEmployee(&req)

	w.WriteHeader(http.StatusAccepted)
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(map[string]interface{}{
		"payload": create,
		"total":   len(create),
	})
}
