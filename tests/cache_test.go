package tests

import (
	"github.com/recrtl/pokeapi-go"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCache(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		getter func() (interface{}, error)
	}{
		{
			name:   "Ability",
			getter: func() (interface{}, error) { return pokeapi.Ability("1") },
		},
		{
			name:   "Berry",
			getter: func() (interface{}, error) { return pokeapi.Berry("1") },
		},
		{
			name:   "BerryFirmness",
			getter: func() (interface{}, error) { return pokeapi.BerryFirmness("1") },
		},
		{
			name:   "BerryFlavor",
			getter: func() (interface{}, error) { return pokeapi.BerryFlavor("1") },
		},
		{
			name:   "Characteristic",
			getter: func() (interface{}, error) { return pokeapi.Characteristic("1") },
		},
		{
			name:   "ContestEffect",
			getter: func() (interface{}, error) { return pokeapi.ContestEffect("1") },
		},
		{
			name:   "ContestType",
			getter: func() (interface{}, error) { return pokeapi.ContestType("1") },
		},
		{
			name:   "EggGroup",
			getter: func() (interface{}, error) { return pokeapi.EggGroup("1") },
		},
		{
			name:   "EncounterCondition",
			getter: func() (interface{}, error) { return pokeapi.EncounterCondition("1") },
		},
		{
			name:   "EncounterConditionValue",
			getter: func() (interface{}, error) { return pokeapi.EncounterConditionValue("1") },
		},
		{
			name:   "EncounterMethod",
			getter: func() (interface{}, error) { return pokeapi.EncounterMethod("1") },
		},
		{
			name:   "EvolutionChain",
			getter: func() (interface{}, error) { return pokeapi.EvolutionChain("1") },
		},
		{
			name:   "EvolutionTrigger",
			getter: func() (interface{}, error) { return pokeapi.EvolutionTrigger("1") },
		},
		{
			name:   "Gender",
			getter: func() (interface{}, error) { return pokeapi.Gender("1") },
		},
		{
			name:   "Generation",
			getter: func() (interface{}, error) { return pokeapi.Generation("1") },
		},
		{
			name:   "GrowthRate",
			getter: func() (interface{}, error) { return pokeapi.GrowthRate("1") },
		},
		{
			name:   "Item",
			getter: func() (interface{}, error) { return pokeapi.Item("1") },
		},
		{
			name:   "ItemAttribute",
			getter: func() (interface{}, error) { return pokeapi.ItemAttribute("1") },
		},
		{
			name:   "ItemCategory",
			getter: func() (interface{}, error) { return pokeapi.ItemCategory("1") },
		},
		{
			name:   "ItemFlingEffect",
			getter: func() (interface{}, error) { return pokeapi.ItemFlingEffect("1") },
		},
		{
			name:   "ItemPocket",
			getter: func() (interface{}, error) { return pokeapi.ItemPocket("1") },
		},
		{
			name:   "Language",
			getter: func() (interface{}, error) { return pokeapi.Language("1") },
		},
		{
			name:   "Location",
			getter: func() (interface{}, error) { return pokeapi.Location("1") },
		},
		{
			name:   "LocationArea",
			getter: func() (interface{}, error) { return pokeapi.LocationArea("1") },
		},
		{
			name:   "Machine",
			getter: func() (interface{}, error) { return pokeapi.Machine("1") },
		},
		{
			name:   "Move",
			getter: func() (interface{}, error) { return pokeapi.Move("1") },
		},
		{
			name:   "MoveAilment",
			getter: func() (interface{}, error) { return pokeapi.MoveAilment("1") },
		},
		{
			name:   "MoveBattleStyle",
			getter: func() (interface{}, error) { return pokeapi.MoveBattleStyle("1") },
		},
		{
			name:   "MoveCategory",
			getter: func() (interface{}, error) { return pokeapi.MoveCategory("1") },
		},
		{
			name:   "MoveDamageClass",
			getter: func() (interface{}, error) { return pokeapi.MoveDamageClass("1") },
		},
		{
			name:   "MoveLearnMethod",
			getter: func() (interface{}, error) { return pokeapi.MoveLearnMethod("1") },
		},
		{
			name:   "MoveTarget",
			getter: func() (interface{}, error) { return pokeapi.MoveTarget("1") },
		},
		{
			name:   "Nature",
			getter: func() (interface{}, error) { return pokeapi.Nature("1") },
		},
		{
			name:   "PalParkArea",
			getter: func() (interface{}, error) { return pokeapi.PalParkArea("1") },
		},
		{
			name:   "PokeathlonStat",
			getter: func() (interface{}, error) { return pokeapi.PokeathlonStat("1") },
		},
		{
			name:   "Pokedex",
			getter: func() (interface{}, error) { return pokeapi.Pokedex("1") },
		},
		{
			name:   "Pokemon",
			getter: func() (interface{}, error) { return pokeapi.Pokemon("1") },
		},
		{
			name:   "PokemonColor",
			getter: func() (interface{}, error) { return pokeapi.PokemonColor("1") },
		},
		{
			name:   "PokemonForm",
			getter: func() (interface{}, error) { return pokeapi.PokemonForm("1") },
		},
		{
			name:   "PokemonHabitat",
			getter: func() (interface{}, error) { return pokeapi.PokemonHabitat("1") },
		},
		{
			name:   "PokemonShape",
			getter: func() (interface{}, error) { return pokeapi.PokemonShape("1") },
		},
		{
			name:   "PokemonSpecies",
			getter: func() (interface{}, error) { return pokeapi.PokemonSpecies("1") },
		},
		{
			name:   "Region",
			getter: func() (interface{}, error) { return pokeapi.Region("1") },
		},
		{
			name:   "Resource",
			getter: func() (interface{}, error) { return pokeapi.Resource("pokemon", 1) },
		},
		{
			name:   "Stat",
			getter: func() (interface{}, error) { return pokeapi.Stat("1") },
		},
		{
			name:   "SuperContestEffect",
			getter: func() (interface{}, error) { return pokeapi.SuperContestEffect("1") },
		},
		{
			name:   "Type",
			getter: func() (interface{}, error) { return pokeapi.Type("1") },
		},
		{
			name:   "Version",
			getter: func() (interface{}, error) { return pokeapi.Version("1") },
		},
		{
			name:   "VersionGroup",
			getter: func() (interface{}, error) { return pokeapi.VersionGroup("1") },
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got1, err1 := tt.getter()
			got2, err2 := tt.getter()
			require.NoError(t, err1)
			require.NoError(t, err2)
			require.Equal(t, got1, got2, "got1 = %v, got2 = %v", got1, got2)
		})
	}
}
