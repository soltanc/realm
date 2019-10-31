package realmdefense

import (
	"crypto/md5"
	"fmt"
)

// v1: var kSecret = []byte{0xE7, 0xA7, 0x98, 0xE5, 0xAF, 0x86}
// v2: var kSecret = []byte{0x8A, 0xA8, 0xF5, 0x92, 0xD7, 0xAC, 0xB4, 0x77, 0x1F, 0x08, 0x90, 0x55, 0x13, 0x23, 0x57, 0x2C}
// v3: var kSecret = []byte{0x29, 0x4e, 0x7a, 0x78, 0x9b, 0x8a, 0xe7, 0xc3, 0xa4, 0xce, 0x24, 0xa8, 0x6b, 0x80, 0x5c, 0x57}
var kSecret = []byte{0x12, 0x81, 0xE4, 0x0C, 0x73, 0xD5, 0x17, 0x77, 0xD8, 0x14, 0x88, 0xD8, 0x0A, 0x70, 0xCF, 0x0E}

// ComputeHash computes the HTTP Header "Hash". Hello, Crypto.ComputeHash ;-)
func ComputeHash(body []byte) string {
	h := md5.New()
	h.Write(body)
	h.Write(kSecret)
	return fmt.Sprintf("%x", h.Sum(nil))
}