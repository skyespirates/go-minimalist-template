package utils

import (
	"math/rand"
	"strings"
)

func mapFn(s string) byte {
	return []byte(s)[0]
}

func MapSlice(in []string, f func(string) byte) []byte {
	result := []byte{}

	for _, val := range in {
		result = append(result, f(val))
	}

	return result
}

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

var alphabets = "abcdefghijklmnopqrstuvwxyz"

func getAlphabets() []byte {

	return MapSlice(strings.Split(alphabets, ""), mapFn)
}

func GenerateKey() string {
	alphabets := getAlphabets()
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
	alphabets := getAlphabets()
	splitKey := MapSlice(strings.Split(key, ""), mapFn)
	splitText := MapSlice(strings.Split(text, ""), mapFn)

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
	alphabets := getAlphabets()
	splitKey := MapSlice(strings.Split(key, ""), mapFn)
	splitEncrypted := MapSlice(strings.Split(encrypted, ""), mapFn)

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
