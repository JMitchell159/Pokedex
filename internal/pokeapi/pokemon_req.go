package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListPokemonInfo(pokemon string) (Pokemon, error) {
	endpoint := "/pokemon/" + pokemon
	fullURL := baseURL + endpoint

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return Pokemon{}, fmt.Errorf("couldn't make the request struct: %v", err)
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, fmt.Errorf("something wrong with the response: %v", err)
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	if res.StatusCode > 399 {
		return Pokemon{}, fmt.Errorf("response failed with status code: %d and \nbody: %s", res.StatusCode, body)
	}

	if err != nil {
		return Pokemon{}, fmt.Errorf("error reading response body: %v", err)
	}

	pokemonDat := Pokemon{}
	err = json.Unmarshal(body, &pokemonDat)
	if err != nil {
		return Pokemon{}, err
	}

	return pokemonDat, nil
}
