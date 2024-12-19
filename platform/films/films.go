package platform

import (
	"encoding/json"
	"io"
	"net/http"
	"sort"
	"sync"
	"time"

	"star-wars-cli/platform"
	planets "star-wars-cli/platform/planets"
	"star-wars-cli/utils"
)

func FilmsListByReleaseDate() error {
	// 1. Get films
	films, err := GetFilmsList()
	if err != nil {
		return err
	}

	// 2. Order by release date
	SortFilms(films)

	// 3. Get the planets information and generate the result
	result := make(map[string][]string)

	// for _, film := range films {
	// 	planetNames, err := planets.GetPlanetByURL(film.Planets)
	// 	if err != nil {
	// 		log.Printf("Error retrieving planets by film %s: %v", film.Title, err)
	// 		continue
	// 	}

	// 	// Sort planet names alphabetically
	// 	sort.Strings(planetNames)
	// 	result[film.Title] = planetNames
	// }

	// using concurrency OPT
	for _, film := range films {
		var wg sync.WaitGroup
		var mu sync.Mutex
		var planetNames []string

		for _, planetURL := range film.Planets {
			wg.Add(1)
			go planets.GetPlanetName(planetURL, &wg, &mu, &planetNames)
		}

		wg.Wait()
		sort.Strings(planetNames)
		result[film.Title] = planetNames
	}

	// 4. Order planets Alpha
	utils.Print(result)

	return nil
}

func GetFilmsList() ([]Films, error) {
	response, err := http.Get(platform.StarWarsAPIUrl + "/films")
	if err != nil {
		return nil, err
	}

	// Read response body
	resBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var films []Films
	err = json.Unmarshal(resBytes, &films)
	if err != nil {
		return nil, err
	}

	return films, nil
}

func SortFilms(films []Films) {
	sort.Slice(films, func(i, j int) bool {

		// Parsing the release date from string to time
		dateObj1, err1 := time.Parse("2006-01-02", films[i].Release_date)
		dateObj2, err2 := time.Parse("2006-01-02", films[j].Release_date)

		// Handle invalid dates
		if err1 != nil || err2 != nil {
			return false
		}

		return dateObj1.Before(dateObj2)
	})
}
