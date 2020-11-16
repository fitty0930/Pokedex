package pokedex

import (
	"github.com/Pokedex/internal/config"
	"github.com/jmoiron/sqlx"
)

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
	db   *sqlx.DB
	conf *config.Config
}

// New ...
func New(db *sqlx.DB, c *config.Config) (PokedexService, error) {
	return service{db, c}, nil
}

func (s service) AddPokemon(m Pokemon) error {
	return nil
}

func (s service) FindByID(ID int) *Pokemon {
	return nil
}

func (s service) FindAll() []*Pokemon {
	var list []*Pokemon
	if err := s.db.Select(&list, "SELECT * FROM pokedex"); err != nil {
		panic(err)
	}
	return list
}
