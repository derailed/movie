package movie_test

import (
	"testing"

	"github.com/derailed/movie"
	"github.com/stretchr/testify/assert"
)

func TestLoad(t *testing.T) {
	data, err := movie.Load("./assets/movies.yml")
	assert.Nil(t, err)

	assert.Equal(t, len(data.Movies), 2)
	assert.Equal(t, "m1", data.Movies[0].Name)
	assert.Equal(t, "ic1", data.Movies[0].Icons[0].Name)
}

func TestLoadFail(t *testing.T) {
	_, err := movie.Load("./assets/blee.yml")
	assert.NotNil(t, err)
}
