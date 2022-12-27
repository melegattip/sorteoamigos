package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

type Controller struct {
	service Service
}

func (c *Controller) AddPerson(w http.ResponseWriter, r *http.Request) {
	var people []Person
	if err := json.NewDecoder(r.Body).Decode(&people); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var wg sync.WaitGroup

	// Add each person to the repository
	for _, person := range people {
		wg.Add(1)
		go func(person Person) {
			defer wg.Done()
			if err := c.service.AddPerson(person); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}(person)
	}

	wg.Wait()

	// Set the custom response header and body
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"message": fmt.Sprintf("%d personas añadidas", len(people)),
	})
}

func (c *Controller) GetGrid(w http.ResponseWriter, r *http.Request) {
	grid, err := c.service.GetGrid()
	if err != nil {
		// Manejar el error y devolver una respuesta adecuada al cliente
		return
	}
	// Devolver la grilla al cliente en formato JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(grid)
}
