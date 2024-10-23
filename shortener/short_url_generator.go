package shortener

import (
	"crypto/sha256"
	"fmt"
	"math/big"
	"os"

	"github.com/itchyny/base58-go"
)

// Generates unique SHA256 hashes for each input string url.
// For more info on implementation, see https://gobyexample.com/sha256-hashes
func sha256HashOf(inputString string) []byte {
	hash := sha256.New()
	hash.Write([]byte(inputString))

	return hash.Sum(nil)
}

// Base58 used here instead of Base64 to omit certain confusing characters and invalid URL characters

// Converts the hashed value to human-readable text
func base58Encode(bytes []byte) string {
	encoding := base58.BitcoinEncoding
	encoded, err := encoding.Encode(bytes)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	return string(encoded)
}

// Main function for generating the shortened link from the encoded original link
func GenerateShortLink(originalLink string, userId string) string {
	// userId concatenated to ensure unique shortened links for the same original link for different users 
	hashedUrl := sha256HashOf(originalLink + userId)
	bit64number := new(big.Int).SetBytes(hashedUrl).Uint64() // converts the 32-byte hash generated from SHA-256 to shorter 64-bit integer
	
	finalString := base58Encode([]byte(fmt.Sprintf("%d", bit64number)))
	
	// Using only first 6 characters keeps the url short and constrained to only 6 characters
	// with a decent level of uniqnuenes with 58^6 possible combinations (from base58 encoding)
	return finalString[:6]
}