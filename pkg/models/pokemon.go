package models

type Pokemon struct {
	Id   float64 `json:"id"`
	Name string  `json:"name"`
}

type RaidPokemon struct {
	Pokemon
	RaidLevel float64 `json:"raid_level"`
}
