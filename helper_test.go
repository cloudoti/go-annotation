package annotation

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Example struct {
	Id string `json:"name=id, len=8"`
}

type Tag struct {
	Name string
	Len  string
}

func TestQuery(t *testing.T) {

	exa := new(Example)
	exa.Id = "1"

	tag := new(Tag)

	err := Parse("json", exa, &tag)

	assert.NoError(t, err)

	assert.Equal(t, tag.Name, "id")
	assert.Equal(t, tag.Len, "8")

}
