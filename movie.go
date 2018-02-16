package movie

import (
	"fmt"
	"io/ioutil"

	"github.com/go-yaml/yaml"
)

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

// Load the movie database
func Load(path string) (Movies, error) {
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
