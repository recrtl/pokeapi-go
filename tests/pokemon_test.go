package tests

import (
	"testing"

	pokeapi "github.com/recrtl/pokeapi-go"
	"github.com/stretchr/testify/assert"
)

func TestAbility(t *testing.T) {
	t.Parallel()
	result, _ := pokeapi.Ability("1")
	assert.Equal(t, "stench", result.Name,
		"Expect to receive Stench.")
}

func TestAbilityByName(t *testing.T) {
	t.Parallel()
	result, _ := pokeapi.Ability("stench")
	assert.Equal(t, "stench", result.Name,
		"Expect to receive Stench.")
}

func TestAbilityFail(t *testing.T) {
	t.Parallel()
	result, _ := pokeapi.Ability("asdf")
	assert.Equal(t, "", result.Name,
		"Expect to receive an empty result.")
}

func TestCharacteristic(t *testing.T) {
	t.Parallel()
	result, _ := pokeapi.Characteristic("1")
	assert.NotZero(t, len(result.Descriptions[1].Description),
		"Expect to receive a description.")
}

func TestCharacteristicFail(t *testing.T) {
	t.Parallel()
	result, _ := pokeapi.Characteristic("asdf")
	assert.Equal(t, 0, len(result.Descriptions),
		"Expect to receive an empty result.")
}

func TestEggGroup(t *testing.T) {
	t.Parallel()
	result, _ := pokeapi.EggGroup("1")
	assert.Equal(t, "monster", result.Name,
		"Expect to receive Monster.")
}

func TestEggGroupByName(t *testing.T) {
	t.Parallel()
	result, _ := pokeapi.EggGroup("monster")
	assert.Equal(t, "monster", result.Name,
		"Expect to receive Monster.")
}

func TestEggGroupFail(t *testing.T) {
	t.Parallel()
	result, _ := pokeapi.EggGroup("asdf")
	assert.Equal(t, "", result.Name,
		"Expect to receive an empty result.")
}

func TestGender(t *testing.T) {
	t.Parallel()
	result, _ := pokeapi.Gender("1")
	assert.Equal(t, "female", result.Name,
		"Expect to receive Female.")
}

func TestGenderByName(t *testing.T) {
	t.Parallel()
	result, _ := pokeapi.Gender("female")
	assert.Equal(t, "female", result.Name,
		"Expect to receive Female.")
}

func TestGenderFail(t *testing.T) {
	t.Parallel()
	result, _ := pokeapi.Gender("asdf")
	assert.Equal(t, "", result.Name,
		"Expect to receive an empty result.")
}

func TestGrowthRate(t *testing.T) {
	t.Parallel()
	result, _ := pokeapi.GrowthRate("1")
	assert.Equal(t, "slow", result.Name,
		"Expect to receive Slow.")
}

func TestGrowthRateByName(t *testing.T) {
	t.Parallel()
	result, _ := pokeapi.GrowthRate("slow")
	assert.Equal(t, "slow", result.Name,
		"Expect to receive Slow.")
}

func TestGrowthRateFail(t *testing.T) {
	t.Parallel()
	result, _ := pokeapi.GrowthRate("asdf")
	assert.Equal(t, "", result.Name,
		"Expect to receive an empty result.")
}

func TestNature(t *testing.T) {
	t.Parallel()
	result, _ := pokeapi.Nature("1")
	assert.Equal(t, "hardy", result.Name,
		"Expect to receive Hardy.")
}

func TestNatureByName(t *testing.T) {
	t.Parallel()
	result, _ := pokeapi.Nature("hardy")
	assert.Equal(t, "hardy", result.Name,
		"Expect to receive Hardy.")
}

func TestNatureFail(t *testing.T) {
	t.Parallel()
	result, _ := pokeapi.Nature("asdf")
	assert.Equal(t, "", result.Name,
		"Expect to receive an empty result.")
}

func TestPokeathlonStat(t *testing.T) {
	t.Parallel()
	result, _ := pokeapi.PokeathlonStat("1")
	assert.Equal(t, "speed", result.Name,
		"Expect to receive Speed.")
}

func TestPokeathlonStatByName(t *testing.T) {
	t.Parallel()
	result, _ := pokeapi.PokeathlonStat("speed")
	assert.Equal(t, "speed", result.Name,
		"Expect to receive Speed.")
}

func TestPokeathlonStatFail(t *testing.T) {
	t.Parallel()
	result, _ := pokeapi.PokeathlonStat("asdf")
	assert.Equal(t, "", result.Name,
		"Expect to receive an empty result.")
}

func TestPokemon(t *testing.T) {
	t.Parallel()
	result, _ := pokeapi.Pokemon("1")
	assert.Equal(t, "bulbasaur", result.Name,
		"Expect to receive Bulbasaur.")
}

func TestPokemonByName(t *testing.T) {
	t.Parallel()
	result, _ := pokeapi.Pokemon("bulbasaur")
	assert.Equal(t, "bulbasaur", result.Name,
		"Expect to receive Bulbasaur.")
}

func TestPokemonFail(t *testing.T) {
	t.Parallel()
	result, _ := pokeapi.Pokemon("digimon")
	assert.Equal(t, "", result.Name,
		"Expect to receive an empty result.")
}

func TestPokemonColor(t *testing.T) {
	t.Parallel()
	result, _ := pokeapi.PokemonColor("1")
	assert.Equal(t, "black", result.Name,
		"Expect to receive Black.")
}

func TestPokemonColorByName(t *testing.T) {
	t.Parallel()
	result, _ := pokeapi.PokemonColor("black")
	assert.Equal(t, "black", result.Name,
		"Expect to receive Black.")
}

func TestPokemonColorFail(t *testing.T) {
	t.Parallel()
	result, _ := pokeapi.PokemonColor("asdf")
	assert.Equal(t, "", result.Name,
		"Expect to receive an empty result.")
}

func TestPokemonForm(t *testing.T) {
	t.Parallel()
	result, _ := pokeapi.PokemonForm("1")
	assert.Equal(t, "bulbasaur", result.Name,
		"Expect to receive Bulbasaur.")
}

func TestPokemonFormByName(t *testing.T) {
	t.Parallel()
	result, _ := pokeapi.PokemonForm("bulbasaur")
	assert.Equal(t, "bulbasaur", result.Name,
		"Expect to receive Bulbasaur.")
}

func TestPokemonFormFail(t *testing.T) {
	t.Parallel()
	result, _ := pokeapi.PokemonForm("asdf")
	assert.Equal(t, "", result.Name,
		"Expect to receive an empty result.")
}

func TestPokemonHabitat(t *testing.T) {
	t.Parallel()
	result, _ := pokeapi.PokemonHabitat("1")
	assert.Equal(t, "cave", result.Name,
		"Expect to receive Cave.")
}

func TestPokemonHabitatByName(t *testing.T) {
	t.Parallel()
	result, _ := pokeapi.PokemonHabitat("cave")
	assert.Equal(t, "cave", result.Name,
		"Expect to receive Cave.")
}

func TestPokemonHabitatFail(t *testing.T) {
	t.Parallel()
	result, _ := pokeapi.PokemonHabitat("asdf")
	assert.Equal(t, "", result.Name,
		"Expect to receive an empty result.")
}

func TestPokemonShape(t *testing.T) {
	t.Parallel()
	result, _ := pokeapi.PokemonShape("1")
	assert.Equal(t, "ball", result.Name,
		"Expect to receive Ball.")
}

func TestPokemonShapeByName(t *testing.T) {
	t.Parallel()
	result, _ := pokeapi.PokemonShape("ball")
	assert.Equal(t, "ball", result.Name,
		"Expect to receive Ball.")
}

func TestPokemonShapeFail(t *testing.T) {
	t.Parallel()
	result, _ := pokeapi.PokemonShape("asdf")
	assert.Equal(t, "", result.Name,
		"Expect to receive an empty result.")
}

func TestPokemonSpecies(t *testing.T) {
	t.Parallel()
	result, _ := pokeapi.PokemonSpecies("1")
	assert.Equal(t, "bulbasaur", result.Name,
		"Expect to receive Bulbasaur.")
}

func TestPokemonSpeciesByName(t *testing.T) {
	t.Parallel()
	result, _ := pokeapi.PokemonSpecies("bulbasaur")
	assert.Equal(t, "bulbasaur", result.Name,
		"Expect to receive Bulbasaur.")
}

func TestPokemonSpeciesFail(t *testing.T) {
	t.Parallel()
	result, _ := pokeapi.PokemonSpecies("asdf")
	assert.Equal(t, "", result.Name,
		"Expect to receive an empty result.")
}

func TestStat(t *testing.T) {
	t.Parallel()
	result, _ := pokeapi.Stat("1")
	assert.Equal(t, "hp", result.Name,
		"Expect to receive HP.")
}

func TestStatByName(t *testing.T) {
	t.Parallel()
	result, _ := pokeapi.Stat("hp")
	assert.Equal(t, "hp", result.Name,
		"Expect to receive HP.")
}

func TestStatFail(t *testing.T) {
	t.Parallel()
	result, _ := pokeapi.Stat("asdf")
	assert.Equal(t, "", result.Name,
		"Expect to receive an empty result.")
}

func TestType(t *testing.T) {
	t.Parallel()
	result, _ := pokeapi.Type("1")
	assert.Equal(t, "normal", result.Name,
		"Expect to receive Normal.")
}

func TestTypeByName(t *testing.T) {
	t.Parallel()
	result, _ := pokeapi.Type("normal")
	assert.Equal(t, "normal", result.Name,
		"Expect to receive Normal.")
}

func TestTypeFail(t *testing.T) {
	t.Parallel()
	result, _ := pokeapi.Type("asdf")
	assert.Equal(t, "", result.Name,
		"Expect to receive an empty result.")
}

func TestPokemonLocationAreas(t *testing.T) {
	t.Parallel()
	result, _ := pokeapi.PokemonLocationAreas("1")
	assert.NotZero(t, len(result))
	assert.Equal(t, "cerulean-city-area", result[0].LocationArea.Name)
	assert.NotZero(t, len(result[0].VersionDetails))
	assert.NotZero(t, len(result[0].VersionDetails[0].EncounterDetails))
	assert.NotZero(t, result[0].VersionDetails[0].EncounterDetails[0].MinLevel)
	assert.NotZero(t, result[0].VersionDetails[0].EncounterDetails[0].MaxLevel)
	assert.NotZero(t, result[0].VersionDetails[0].EncounterDetails[0].Chance)
}
