package main

import (
	"math/rand"
	"strings"
	"testing"
	"time"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

var basicTestFixtures = []struct {
	str string
	min int
}{
	{"bitcoin take over the world maybe who knows perhaps", 3},
	{"turns out random test cases are easier than writing out basic ones", 3},
	{"lets talk about javascript the best language", 3},
	{"i want to travel the world writing code one day", 1},
	{"Lets all go on holiday somewhere very cold", 2},
}

// letters needs for generating random strings
const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func TestFindShortBasic(t *testing.T) {
	for _, tc := range basicTestFixtures {
		res := FindShort(tc.str)
		if tc.min != tc.min {
			t.Errorf("text:%v: expected %v; got %v", tc.str, tc.min, res)
		}
	}
}

func TestFindShortRandom(t *testing.T) {
	for i := 0; i < 1000; i++ {
		minWordLength := randInt(1, 10)
		maxWordLength := randInt(minWordLength, 64)
		wordsNum := randInt(1, 32)

		str := randStringOfWords(wordsNum, minWordLength, maxWordLength)
		res := FindShort(str)
		if res != minWordLength {
			t.Errorf("text:%v: expected %v; got %v", str, minWordLength, res)
		}
	}
}

func randWord(length int) string {
	str := make([]byte, length)
	for i := range str {
		str[i] = letters[rand.Intn(len(letters))]
	}

	return string(str)
}

func randStringOfWords(wordNum, minWordLength, maxWordLength int) string {
	words := make([]string, wordNum)
	// generate shortest word
	words[0] = randWord(minWordLength)
	// generate other words
	for i := 1; i < wordNum; i++ {
		wordLength := randInt(minWordLength, maxWordLength)
		words[i] = randWord(wordLength)
	}
	rand.Shuffle(len(words), func(i, j int) {
		words[i], words[j] = words[j], words[i]
	})

	return strings.Join(words, " ")
}

func randInt(from, to int) int {
	return rand.Intn(to-from+1) + from
}
