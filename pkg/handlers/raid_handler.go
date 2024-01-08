package handlers

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/Israel-Ferreira/servidor-em-go/pkg/models"
)

const POGO_API = "https://pogoapi.net/api/v1"

func InternalErrorResponse(rw http.ResponseWriter, msg string) {
	msg_resp := map[string]string{
		"msg": msg,
	}

	rw.WriteHeader(http.StatusInternalServerError)

	json.NewEncoder(rw).Encode(&msg_resp)
}



func cleanResult(body []byte) ([]models.RaidPokemon, error) {

	var result map[string]any

	var pokemons []models.RaidPokemon

	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}


	for _, pokemon := range result {

		dict := pokemon.(map[string]any)

		pokemonId := dict["id"].(float64)

		raidPokemon := models.RaidPokemon{
			Pokemon: models.Pokemon{
				Id:   pokemonId,
				Name: dict["name"].(string),
			},
			RaidLevel: dict["raid_level"].(float64),
		}

		pokemons = append(pokemons, raidPokemon)
	}

	return pokemons, nil
}



func GetRaidExclusivePokemons(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		log.Println("ERRO: Método HTTP não Permitido para essa rota")
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	urlRaidPokemons := fmt.Sprintf("%s/raid_exclusive_pokemon.json", POGO_API)

	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	resp, err := http.Get(urlRaidPokemons)

	if err != nil {
		log.Println("ERRO: ", err)
		InternalErrorResponse(w, "Erro ao buscar os dados da API")

		return
	}

	log.Println(resp.StatusCode)

	defer resp.Body.Close()

	jsonBody, err := io.ReadAll(resp.Body)

	if err != nil {
		InternalErrorResponse(w, "Erro ao fazer o parse da resposta")
		return
	}

	cleanedResult, err := cleanResult(jsonBody)

	if err != nil {
		msg := map[string]string{
			"msg": "Erro ao fazer o parse da resposta",
		}

		w.WriteHeader(http.StatusInternalServerError)

		json.NewEncoder(w).Encode(&msg)

		return
	}

	if err = json.NewEncoder(w).Encode(&cleanedResult); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}
