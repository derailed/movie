package movie_test

import (
	"testing"

	"github.com/derailed/nuclio/movie"
	"github.com/stretchr/testify/assert"
)

func TestLoadFile(t *testing.T) {
	data, err := movie.LoadFile("./assets/movies.yml")
	assert.Nil(t, err)

	assert.Equal(t, len(data.Movies), 2)
	assert.Equal(t, "m1", data.Movies[0].Name)
	assert.Equal(t, "ic1", data.Movies[0].Icons[0].Name)
}

func TestLoadFileFail(t *testing.T) {
	_, err := movie.LoadFile("./assets/blee.yml")
	assert.NotNil(t, err)
}

func TestLoadMem(t *testing.T) {
	data, err := movie.LoadMem()
	assert.Nil(t, err)

	assert.Equal(t, len(data.Movies), 3)
	assert.Equal(t, "Home Alone", data.Movies[0].Name)
	assert.Equal(t, "🏡", data.Movies[0].Icons[0].Name)
}
