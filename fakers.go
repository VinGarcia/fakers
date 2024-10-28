package fakers

import (
	"github.com/vingarcia/structscanner"
)

func New(obj any, customValues map[string]any) error {
	return structscanner.Decode(obj, newCustomValuesMapDecoder(customValues))
}
