// Package json5 support for parse and load json5
package json5

import (
	"github.com/gookit/config/v2"
	"github.com/yosuke-furukawa/json5/encoding/json5"
)

// NAME for driver
const NAME = "json5"

// Name for driver
const Name = "json5"

// JSONMarshalIndent if not empty, will use json.MarshalIndent for encode data.
var JSONMarshalIndent string

var (
	// Decoder for json
	Decoder config.Decoder = json5.Unmarshal

	// Encoder for json5
	Encoder config.Encoder = func(v any) (out []byte, err error) {
		if len(JSONMarshalIndent) == 0 {
			return json5.Marshal(v)
		}

		return json5.MarshalIndent(v, "", JSONMarshalIndent)
	}

	// Driver for json5
	Driver = config.NewDriver(Name, Decoder, Encoder)
)
