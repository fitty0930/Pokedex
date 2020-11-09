package main

import "testing"

func TestPokedexAdd(t *testing.T) {
	p := NewPokedex()
	m1 := p.FindByID(1)
	if m1 != nil {
		t.Error("El pokemon con el ID 1 ya existe")
	}
	p.Add(Pokemon{1, "Bulbasaur"})
	m1 = p.FindByID(1)
	if m1 == nil {
		t.Error("El pokemon con el ID 1 no se agrego")
	}

	if m1.Name != "Bulbasaur" {
		t.Error("El pokemon con el ID 1 no tiene el nombre correcto")
	}
}
