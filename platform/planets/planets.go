package platform

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"

	"star-wars-cli/platform"
	"star-wars-cli/utils"
)

func GetPlanetByURL(planetURLs []string) ([]string, error) {

	var planetNames []string

	for _, planetURL := range planetURLs {
		response, err := http.Get(planetURL)
		if err != nil {
			return nil, err
		}

		// Read response body
		resBytes, err := io.ReadAll(response.Body)
		if err != nil {
			return nil, err
		}

		var planet Planet
		err = json.Unmarshal(resBytes, &planet)
		if err != nil {
			return nil, err
		}

		planetNames = append(planetNames, planet.Name)
	}

	return planetNames, nil
}

func GetPlanetsList() error {

	response, err := http.Get(platform.StarWarsAPIUrl + "/planets")
	if err != nil {
		return err
	}

	// Read response body
	resBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}

	var planets []Planet
	err = json.Unmarshal(resBytes, &planets)
	if err != nil {
		return err
	}

	utils.Print(planets)

	return nil
}

func GetPlanetName(url string, wg *sync.WaitGroup, mu *sync.Mutex, planetNames *[]string) {
	defer wg.Done()

	var planet Planet
	response, err := http.Get(url)
	if err != nil {
		fmt.Printf("Failed getting planet: %v\n", err)
		return
	}

	// Read response body
	resBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return
	}

	err = json.Unmarshal(resBytes, &planet)
	if err != nil {
		return
	}

	mu.Lock()
	*planetNames = append(*planetNames, planet.Name)
	mu.Unlock()
}
