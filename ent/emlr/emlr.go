// package emlr contains a reader for the EML file format.
package emlr

import (
	"bytes"
	"encoding/xml"
	"io"
	"log/slog"

	"github.com/gnames/ecoml/ent/eml"
)

type emlr struct {
}

// New returns a new EML reader.
func New() EMLReader {
	return &emlr{}
}

func (e *emlr) Decode(r io.Reader) (*eml.EML, error) {
	bs, err := io.ReadAll(r)
	if err != nil {
		slog.Error("Cannot read EML file", "error", err)
		return nil, err
	}
	var res eml.EML
	decoder := xml.NewDecoder(bytes.NewReader(bs))
	err = decoder.Decode(&res)
	if err != nil {
		slog.Error("Cannot decode EML file", "error", err)
		return nil, err
	}
	return &res, nil
}
