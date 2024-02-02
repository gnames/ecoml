package emlr_test

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/gnames/ecoml/ent/emlr"
	"github.com/stretchr/testify/assert"
)

func TestDecode(t *testing.T) {
	assert := assert.New(t)
	path := filepath.Join("..", "..", "testdata", "simple.xml")
	f, err := os.Open(path)
	assert.Nil(err)
	defer f.Close()

	e := emlr.New()
	eml, err := e.Decode(f)
	assert.Nil(err)
	fmt.Printf("EML: %#v\n", eml)
}
