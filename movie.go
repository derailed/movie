package movie

import (
	"fmt"
	"io/ioutil"

	"github.com/go-yaml/yaml"
)

// MOVIES in memory movie corpus
const MOVIES = `
movies:
- name: m1
  icons:
  - name: ic1
  - name: ic2
- name: m2
  icons:
  - name: ic11
	- name: ic21
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
func LoadMem(path string) (Movies, error) {
	var movies Movies

	err = yaml.Unmarshal(data, &movies)
	if err != nil {
		return movies, err
	}

	return movies, nil
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
		fmt.Println("YO!")
		return movies, err
	}

	return movies, nil
}
