package structs

type EvolutionDetail struct {
	Gender                int         `json:"gender"`
	HeldItem              interface{} `json:"held_item"`
	Item                  interface{} `json:"item"`
	KnownMove             interface{} `json:"known_move"`
	KnownMoveType         interface{} `json:"known_move_type"`
	Location              interface{} `json:"location"`
	MinAffection          int         `json:"min_affection"`
	MinBeauty             int         `json:"min_beauty"`
	MinHappiness          int         `json:"min_happiness"`
	MinLevel              int         `json:"min_level"`
	NeedsOverworldRain    bool        `json:"needs_overworld_rain"`
	PartySpecies          interface{} `json:"party_species"`
	PartyType             interface{} `json:"party_type"`
	RelativePhysicalStats int         `json:"relative_physical_stats"`
	TimeOfDay             string      `json:"time_of_day"`
	TradeSpecies          interface{} `json:"trade_species"`
	Trigger               struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"trigger"`
	TurnUpsideDown bool `json:"turn_upside_down"`
}

type EvolutionChainLink struct {
	EvolutionDetails []EvolutionDetail    `json:"evolution_details"`
	EvolvesTo        []EvolutionChainLink `json:"evolves_to"`
	IsBaby           bool                 `json:"is_baby"`
	Species          struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"species"`
}

// EvolutionChain is a single evolution chain.
type EvolutionChain struct {
	BabyTriggerItem interface{}        `json:"baby_trigger_item"`
	Chain           EvolutionChainLink `json:"chain"`
	ID              int                `json:"id"`
}

// EvolutionTrigger is a single evolution trigger.
type EvolutionTrigger struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PokemonSpecies []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"pokemon_species"`
}
