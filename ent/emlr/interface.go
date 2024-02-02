package emlr

import (
	"io"

	"github.com/gnames/ecoml/ent/eml"
)

type EMLReader interface {
	Decode(io.Reader) (*eml.EML, error)
}
