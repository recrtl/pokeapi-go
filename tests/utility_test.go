package tests

import (
	"testing"

	pokeapi "github.com/recrtl/pokeapi-go"
	"github.com/stretchr/testify/assert"
)

func TestLanguage(t *testing.T) {
	t.Parallel()
	result, _ := pokeapi.Language("1")
	assert.Equal(t, "ja-Hrkt", result.Name,
		"Expect to receive Japanese.")
}

func TestLanguageByName(t *testing.T) {
	t.Parallel()
	result, _ := pokeapi.Language("ja-Hrkt")
	assert.Equal(t, "ja-Hrkt", result.Name,
		"Expect to receive Japanese.")
}

func TestLanguageFail(t *testing.T) {
	t.Parallel()
	result, _ := pokeapi.Language("asdf")
	assert.Equal(t, "", result.Name,
		"Expect to receive an empty result.")
}
