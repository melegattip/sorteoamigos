package main

type Repository interface {
	AddPerson(p Person) error
	GetPeople() []Person
}

type InMemoryRepository struct {
	people []Person
}

func (r *InMemoryRepository) AddPerson(p Person) error {
	r.people = append(r.people, p)
	return nil
}

func (r *InMemoryRepository) GetPeople() []Person {
	return r.people
}
