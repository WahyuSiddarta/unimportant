package helper

import (
	"github.com/minio/sha256-simd"
)

//SHA256b :
func SHA256b(msg []byte) []byte {
	x := sha256.Sum256(msg)
	return x[:]
}
