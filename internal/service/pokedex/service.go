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

// Service ...
type Service interface {
	AddPokemon(string) error
	FindByID(int64) *Pokemon
	FindAll() []*Pokemon
	DeleteByID(int64) error
	ChangePokemon(ID int64, m string) error
}

type service struct {
	db   *sqlx.DB
	conf *config.Config
}

// New ...
func New(db *sqlx.DB, c *config.Config) (Service, error) {
	return service{db, c}, nil
}

func (s service) AddPokemon(m string) error {
	s.db.MustExec("INSERT INTO pokedex (name) VALUES (?)", m)
	return nil
}

func (s service) ChangePokemon(ID int64, m string) error {
	s.db.MustExec("UPDATE pokedex SET Name=? WHERE ID=?", m, ID)
	return nil
}

func (s service) DeleteByID(ID int64) error {
	err := s.db.MustExec("DELETE FROM pokedex WHERE ID = ?", ID)
	if err != nil {
		return nil
	}
	return nil
}

func (s service) FindByID(ID int64) *Pokemon {
	// mia
	var pokemon Pokemon
	err := s.db.QueryRow("SELECT * FROM pokedex WHERE ID = ?", ID).Scan(&pokemon.ID, &pokemon.Name)
	if err != nil {
		return nil
	}
	return &pokemon
}

func (s service) FindAll() []*Pokemon {
	var list []*Pokemon
	if err := s.db.Select(&list, "SELECT * FROM pokedex"); err != nil {
		panic(err)
	}
	return list
}
