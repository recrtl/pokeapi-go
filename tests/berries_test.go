package tests

import (
	"testing"

	pokeapi "github.com/recrtl/pokeapi-go"
	"github.com/stretchr/testify/assert"
)

func TestBerry(t *testing.T) {
	t.Parallel()
	result, _ := pokeapi.Berry("1")
	assert.Equal(t, "cheri", result.Name,
		"Expect to receive Cheri.")
}

func TestBerryByName(t *testing.T) {
	t.Parallel()
	result, _ := pokeapi.Berry("cheri")
	assert.Equal(t, "cheri", result.Name,
		"Expect to receive Cheri.")
}

func TestBerryFail(t *testing.T) {
	t.Parallel()
	result, _ := pokeapi.Berry("asdf")
	assert.Equal(t, "", result.Name,
		"Expect to receive an empty result.")
}

func TestBerryFirmness(t *testing.T) {
	t.Parallel()
	result, _ := pokeapi.BerryFirmness("1")
	assert.Equal(t, "very-soft", result.Name,
		"Expect to receive Very Soft.")
}

func TestBerryFirmnessByName(t *testing.T) {
	t.Parallel()
	result, _ := pokeapi.BerryFirmness("very-soft")
	assert.Equal(t, "very-soft", result.Name,
		"Expect to receive Very Soft.")
}

func TestBerryFirmnessFail(t *testing.T) {
	t.Parallel()
	result, _ := pokeapi.BerryFirmness("asdf")
	assert.Equal(t, "", result.Name,
		"Expect to receive an empty result.")
}

func TestBerryFlavor(t *testing.T) {
	t.Parallel()
	result, _ := pokeapi.BerryFlavor("1")
	assert.Equal(t, "spicy", result.Name,
		"Expect to receive Spicy.")
}

func TestBerryFlavorByName(t *testing.T) {
	t.Parallel()
	result, _ := pokeapi.BerryFlavor("spicy")
	assert.Equal(t, "spicy", result.Name,
		"Expect to receive Spicy.")
}

func TestBerryFlavorFail(t *testing.T) {
	t.Parallel()
	result, _ := pokeapi.BerryFlavor("asdf")
	assert.Equal(t, "", result.Name,
		"Expect to receive an empty result.")
}
