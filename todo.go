package main

import "time"

// Todo main record structure
type Todo struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Completed bool      `json:"completed"`
	Due       time.Time `json:"due"`
}

// Todos array of Todo
type Todos []Todo

// GetTodos return todos records
func GetTodos() Todos {
	todos := Todos{
		Todo{ID: 0, Name: "Write presentation", Completed: true},
		Todo{ID: 1, Name: "Host meetup", Completed: true},
		Todo{ID: 2, Name: "Make something real !!!", Completed: false},
	}
	return todos
}
