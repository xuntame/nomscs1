package random

import (
	"crypto/rand"
	"encoding/hex"

	"github.com/stormasm/noms/go/d"
)

var (
	reader = rand.Reader
)

// Creates a unique ID which is a random 16 byte hex string
func Id() string {
	data := make([]byte, 16)
	_, err := reader.Read(data)
	d.Chk.NoError(err)
	return hex.EncodeToString(data)
}
