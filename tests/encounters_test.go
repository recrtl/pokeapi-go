package tests

import (
	"testing"

	pokeapi "github.com/recrtl/pokeapi-go"
	"github.com/stretchr/testify/assert"
)

func TestEncounterMethod(t *testing.T) {
	t.Parallel()
	result, _ := pokeapi.EncounterMethod("1")
	assert.Equal(t, "walk", result.Name,
		"Expect to receive Walk.")
}

func TestEncounterMethodByName(t *testing.T) {
	t.Parallel()
	result, _ := pokeapi.EncounterMethod("walk")
	assert.Equal(t, "walk", result.Name,
		"Expect to receive Walk.")
}

func TestEncounterMethodFail(t *testing.T) {
	t.Parallel()
	result, _ := pokeapi.EncounterMethod("asdf")
	assert.Equal(t, "", result.Name,
		"Expect to receive an empty result.")
}

func TestEncounterCondition(t *testing.T) {
	t.Parallel()
	result, _ := pokeapi.EncounterCondition("1")
	assert.Equal(t, "swarm", result.Name,
		"Expect to receive Swarm.")
}

func TestEncounterConditionByName(t *testing.T) {
	t.Parallel()
	result, _ := pokeapi.EncounterCondition("swarm")
	assert.Equal(t, "swarm", result.Name,
		"Expect to receive Swarm.")
}

func TestEncounterConditionFail(t *testing.T) {
	t.Parallel()
	result, _ := pokeapi.EncounterCondition("asdf")
	assert.Equal(t, "", result.Name,
		"Expect to receive an empty result.")
}

func TestEncounterConditionValue(t *testing.T) {
	t.Parallel()
	result, _ := pokeapi.EncounterConditionValue("1")
	assert.Equal(t, "swarm-yes", result.Name,
		"Expect to receive Swarm (yes).")
}

func TestEncounterConditionValueByName(t *testing.T) {
	t.Parallel()
	result, _ := pokeapi.EncounterConditionValue("swarm-yes")
	assert.Equal(t, "swarm-yes", result.Name,
		"Expect to receive Swarm (yes).")
}

func TestEncounterConditionValueFail(t *testing.T) {
	t.Parallel()
	result, _ := pokeapi.EncounterConditionValue("asdf")
	assert.Equal(t, "", result.Name,
		"Expect to receive an empty result.")
}
