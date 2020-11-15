package pokedex

import "github.com/Pokedex/internal/config"

// Pokemon ... como se define cada pokemon en la pokedex
type Pokemon struct {
	ID   int64
	Name string
}

// PokedexService ...
type PokedexService interface {
	AddPokemon(Pokemon) error
	FindByID(int) *Pokemon
	FindAll() []*Pokemon
}

type service struct {
	conf *config.Config
}

// New ...
func New(c *config.Config) (PokedexService, error) {
	return service{c}, nil
}

func (s service) AddPokemon(m Pokemon) error {
	return nil
}

func (s service) FindByID(ID int) *Pokemon {
	return nil
}

func (s service) FindAll() []*Pokemon {
	var list []*Pokemon
	list = append(list, &Pokemon{1, "Bulbasaur"})
	return list
}
