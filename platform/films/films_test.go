package platform

import (
	"encoding/json"
	"net/http"
	"reflect"
	"star-wars-cli/platform"
	"testing"

	"github.com/h2non/gock"
)

func TestSortFilms(t *testing.T) {
	type args struct {
		films []Films
	}
	tests := []struct {
		name     string
		args     args
		expected []Films
	}{
		{
			name: "C01: ",
			args: args{
				films: []Films{
					{
						Title:        "second",
						Release_date: "2016-01-28",
						Planets:      []string{"b", "c", "a"},
					},
					{
						Title:        "first",
						Release_date: "2014-01-28",
						Planets:      []string{"b", "a", "c"},
					},
				},
			},
			expected: []Films{
				{
					Title:        "first",
					Release_date: "2014-01-28",
					Planets:      []string{"b", "a", "c"},
				},
				{
					Title:        "second",
					Release_date: "2016-01-28",
					Planets:      []string{"b", "c", "a"},
				},
			},
		},
		{
			name: "C02: Invalid dates",
			args: args{
				films: []Films{
					{
						Title:        "second",
						Release_date: "201604-01-28",
						Planets:      []string{"b", "c", "a"},
					},
					{
						Title:        "first",
						Release_date: "201449-01-28",
						Planets:      []string{"b", "a", "c"},
					},
				},
			},
			expected: []Films{
				{
					Title:        "second",
					Release_date: "201604-01-28",
					Planets:      []string{"b", "c", "a"},
				},
				{
					Title:        "first",
					Release_date: "201449-01-28",
					Planets:      []string{"b", "a", "c"},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SortFilms(tt.args.films)

			// compare expected output
			isExpectedOut := CompareSlices(tt.args.films, tt.expected)
			if !isExpectedOut {
				t.Errorf("isExpectedOut = true, output %v", isExpectedOut)
				return
			}
		})
	}
}

// CompareSlices compares two slices and returns true if they are identical.
func CompareSlices(slice1, slice2 []Films) bool {
	if len(slice1) != len(slice2) {
		return false
	}
	for i := 0; i < len(slice1); i++ {
		if slice1[i].Title != slice2[i].Title {
			return false
		}
	}
	return true
}

func TestGetFilmsList(t *testing.T) {

	//C01 mock response model
	mockRes := []Films{
		{
			Title:        "Film1",
			Release_date: "2012-01-01",
			Planets:      []string{},
		},
	}

	mockBytes, _ := json.Marshal(mockRes)

	tests := []struct {
		name     string
		want     []Films
		wantErr  bool
		mockData func()
	}{
		{
			name: "C01: GetFilmsList: STATUS OK",
			want: []Films{
				{
					Title:        "Film1",
					Release_date: "2012-01-01",
					Planets:      []string{},
				},
			},
			wantErr: false,
			mockData: func() {
				// Mock server endpoint
				platform.StarWarsAPIUrl = "http://swapi.info/api"
				gock.New(platform.StarWarsAPIUrl).
					Get("/films").
					Reply(http.StatusOK).
					BodyString(string(mockBytes))
			},
		},
		{
			name:    "C02: GetFilmsList: Err",
			want:    nil,
			wantErr: true,
			mockData: func() {
				// Mock server endpoint
				platform.StarWarsAPIUrl = "http://swapi.info/api"
				gock.New(platform.StarWarsAPIUrl).
					Get("/films/mock").
					Reply(http.StatusOK).
					BodyString(string(mockBytes))
			},
		},
		{
			name:    "C03: GetFilmsList: Err on json.Unmarshal",
			want:    nil,
			wantErr: true,
			mockData: func() {
				// Mock server endpoint
				platform.StarWarsAPIUrl = "http://swapi.info/api"
				gock.New(platform.StarWarsAPIUrl).
					Get("/films").
					Reply(http.StatusOK).
					BodyString("test")
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockData()
			got, err := GetFilmsList()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetFilmsList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetFilmsList() = %v, want %v", got, tt.want)
			}
		})
	}
}


