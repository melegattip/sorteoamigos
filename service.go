package main

import "math/rand"

type Service interface {
	AddPerson(p Person) error
	GetGrid() ([]response, error)
}

type response struct {
	Entrega, Recibe string
}

type PersonService struct {
	repository Repository
}

func (s *PersonService) AddPerson(p Person) error {
	return s.repository.AddPerson(p)
}

func (s *PersonService) GetGrid() ([]response, error) {
	people := s.repository.GetPeople()
	return generateGrid(people), nil
}

func generateGrid(people []Person) []response {
	// Crear una lista para almacenar las tuplas de personas
	grid := make([]response, 0)
	// Desordenar aleatoriamente la lista de personas
	rand.Shuffle(len(people), func(i, j int) { people[i], people[j] = people[j], people[i] })
	// Recorrer la lista de personas y generar las tuplas de personas
	for i := 0; i < len(people); i++ {
		p1 := people[i]
		p2 := people[(i+1)%len(people)]
		result := response{p1.Name, p2.Name}
		grid = append(grid, result)
	}
	// Devolver la grilla
	return grid
}
