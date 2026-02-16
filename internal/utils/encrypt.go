package utils

import (
	"math/rand"
	"strings"
)

func findMatch(dictionary map[byte]byte, in byte) byte {

	var char byte
	v, ok := dictionary[in]

	if !ok {
		if in >= 65 && in <= 90 {
			lower := in + 32           // to lowercase
			match := dictionary[lower] // find matched key
			upper := match - 32        // to uppercase
			char = upper
		} else {
			char = in
		}
	} else {
		char = v
	}

	return char

}

const chars = "abcdefghijklmnopqrstuvwxyz"

func GenerateKey() string {
	alphabets := []byte(chars)
	stringKeys := make([]string, len(alphabets))
	for i, val := range alphabets {
		stringKeys[i] = string(val)
	}

	rand.Shuffle(len(stringKeys), func(i, j int) {
		stringKeys[i], stringKeys[j] = stringKeys[j], stringKeys[i]
	})

	key := strings.Join(stringKeys, "")
	return key
}

func Encrypt(key, text string) string {
	alphabets := []byte(chars)
	splitKey := []byte(key)
	splitText := []byte(text)

	result := []byte{}
	dictionary := make(map[byte]byte)

	for i, val := range alphabets {
		dictionary[val] = splitKey[i]
	}

	for _, val := range splitText {
		char := findMatch(dictionary, val)

		result = append(result, char)
	}
	return string(result)
}

func Decrypt(key, encrypted string) string {
	alphabets := []byte(chars)
	splitKey := []byte(key)
	splitEncrypted := []byte(encrypted)

	dictionary := map[byte]byte{}

	for i, val := range splitKey {
		dictionary[val] = alphabets[i]
	}

	result := []byte{}

	for _, val := range splitEncrypted {
		char := findMatch(dictionary, val)

		result = append(result, char)
	}

	return string(result)
}
