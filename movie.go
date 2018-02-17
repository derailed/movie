package movie

import (
	"io/ioutil"

	"github.com/go-yaml/yaml"
)

// MOVIES in memory movie corpus
const MOVIES = `
movies:
- name: Home Alone
  icons:
  - name: ! "🏡"
  - name: ! "🍣"
- name: Terminator
  icons:
  - name: ! "🏡"
  - name: ! "👻"
- name: Ghost
  icons:
  - name: ! "🏡"
  - name: ! "👻"
`

type (
	// Icon an emoji representation
	Icon struct {
		Name string `yaml:"name"`
	}

	// Movie an Iconoflix movie representation
	Movie struct {
		Name  string `yaml:"name"`
		Icons []Icon `yaml':",inline"`
	}

	// Movies a collection of Iconoflix movies
	Movies struct {
		Movies []Movie `yaml:"movies"`
	}
)

// LoadMem the movie database from memory
func LoadMem() (d Movies, err error) {
	err = yaml.Unmarshal([]byte(MOVIES), &d)
	return
}

// LoadFile the movie database from file
func LoadFile(path string) (Movies, error) {
	var movies Movies

	data, err := ioutil.ReadFile(path)
	if err != nil {
		return movies, err
	}

	err = yaml.Unmarshal(data, &movies)
	if err != nil {
		return movies, err
	}

	return movies, nil
}
