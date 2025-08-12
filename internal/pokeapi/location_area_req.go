package pokeapi

import (
	"encoding/json"
	"net/http"
	"fmt"
	"io"
)

func (c *Client) ListLocationAreas(pageURL *string) (LocationAreasResp, error) {
	endpoint := "/location-area"
	fullURL := baseURL + endpoint
	if pageURL != nil {
		fullURL = *pageURL
	}

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return LocationAreasResp{}, fmt.Errorf("couldn't make the request struct: %v", err)
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreasResp{}, fmt.Errorf("something wrong with the response: %v", err)
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	if res.StatusCode > 399 {
		return LocationAreasResp{}, fmt.Errorf("response failed with status code: %d and \nbody: %s", res.StatusCode, body)
	}

	if err != nil {
		return LocationAreasResp{}, fmt.Errorf("error reading response body: %v", err)
	}

	var locationAreasResp LocationAreasResp
	err = json.Unmarshal(body, &locationAreasResp)
	if err != nil {
		return LocationAreasResp{}, err
	}

	return locationAreasResp, nil
}