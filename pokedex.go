package main

import "fmt"

// Pokedex ...
type Pokedex struct {
	pokemones map[int]*Pokemon
}

// Pokemon ...
type Pokemon struct {
	ID   int
	Name string
}

// NewPokedex ...
func NewPokedex() Pokedex {
	pokemones := make(map[int]*Pokemon)
	return Pokedex{
		pokemones,
	}
}

// Add ...
func (p Pokedex) Add(m Pokemon) {
	// m por monster
	p.pokemones[m.ID] = &m
}

// Print ...
func (p Pokedex) Print() {
	for _, v := range p.pokemones {
		fmt.Printf("[%v]\t%v\n", v.ID, v.Name)
	}
}

// FindByID ...
func (p Pokedex) FindByID(ID int) *Pokemon {
	return p.pokemones[ID]
}

// Delete ...
func (p Pokedex) Delete(ID int) {
	delete(p.pokemones, ID)
}

// Update ...
func (p Pokedex) Update(m Pokemon) {
	p.pokemones[m.ID] = &m
}
