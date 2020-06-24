package parameters

import (
	"crypto/rand"
	"encoding/hex"
	"io"
)

// Generate a random name for attachment.
func randomAttachName() string {
	buf := make([]byte, 10)
	_, _ = io.ReadFull(rand.Reader, buf)
	return hex.EncodeToString(buf)
}
