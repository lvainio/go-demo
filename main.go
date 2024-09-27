package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func main() {
	text := "Secure SHA-256 Hash"

	hash := sha256.New()
	hash.Write([]byte(text))

	checksum := hex.EncodeToString(hash.Sum(nil))

	fmt.Printf("SHA-256 hash of '%s' is: %s\n", text, checksum)
}
