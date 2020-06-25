package parameters

import (
	"crypto/rand"
	"encoding/hex"
	"io"
)

const schemeAttach = "attach://"

// Generate a random name for attachment.
func randomAttachNameAndUri() (name string, uri string) {
	buf := make([]byte, 10)
	_, _ = io.ReadFull(rand.Reader, buf)
	name = hex.EncodeToString(buf)
	uri = schemeAttach + name
	return
}
