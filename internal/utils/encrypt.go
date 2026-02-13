package utils

import (
	"math/rand"
	"strings"
)

// map []string to []rune
func mapStringToRune(in []string, f func(s string) rune) []rune {

	result := []rune{}

	for _, val := range in {
		result = append(result, f(val))
	}

	return result

}

// map []rune to []string
func mapRuneToString(in []rune, f func(r rune) string) []string {

	result := []string{}

	for _, val := range in {
		result = append(result, f(val))
	}

	return result

}

func findMatch(dictionary map[rune]rune, in rune) rune {

	var char rune
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

func getAlphabets() []rune {
	alphabets := make([]rune, 0)
	for i := 97; i <= 122; i++ {
		char := rune(i)
		alphabets = append(alphabets, char)
	}

	return alphabets
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
	splitKey := mapStringToRune(strings.Split(key, ""), func(s string) rune { return []rune(s)[0] })
	splitText := mapStringToRune(strings.Split(text, ""), func(s string) rune { return []rune(s)[0] })

	result := []rune{}
	dictionary := make(map[rune]rune)

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
	splitKey := mapStringToRune(strings.Split(key, ""), func(s string) rune { return []rune(s)[0] })
	splitEncrypted := mapStringToRune(strings.Split(encrypted, ""), func(s string) rune { return []rune(s)[0] })

	dictionary := make(map[rune]rune)

	for i, val := range splitKey {
		dictionary[val] = alphabets[i]
	}

	result := []rune{}

	for _, val := range splitEncrypted {
		char := findMatch(dictionary, val)

		result = append(result, char)
	}

	r := mapRuneToString(result, func(r rune) string { return string(r) })

	return strings.Join(r, "")
}
