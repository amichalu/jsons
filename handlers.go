package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Index / handler
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

// TodoIndex todos/ handler
func TodoIndex(w http.ResponseWriter, r *http.Request) {
	todos := GetTodos()

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(todos); err != nil {
		panic(err)
	}
}

// TodoShow todo
func TodoShow(w http.ResponseWriter, r *http.Request) {

	// Get ID from request
	vars := mux.Vars(r)
	todoIDi, err := strconv.Atoi(vars["todoId"])
	if err != nil {
		fmt.Fprintln(w, "Nothing to do for ID:", todoIDi)
		return
	}

	// Get data and searching searching
	todos := GetTodos()
	var found = false
	var result = Todo{}
	for i, v := range todos {
		if v.ID == todoIDi {
			found = true
			result = todos[i]
			break
		}
	}

	// Outputing
	if found {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(result); err != nil {
			panic(err)
		}
	} else {
		fmt.Fprintln(w, "Nothing to do for ID:", todoIDi)
	}

}
