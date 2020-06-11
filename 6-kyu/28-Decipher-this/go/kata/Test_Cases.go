package kata_test

import (
	. "codewarrior/kata"
	"math/rand"
	"strconv"
	"strings"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Test Suite", func() {
	It("Fixed Tests", func() {
		Expect(DecipherThis("")).Should(Equal(""))
		Expect(DecipherThis("65")).Should(Equal("A"))
		Expect(DecipherThis("84eh")).Should(Equal("The"))
		Expect(DecipherThis("65 66 67")).Should(Equal("A B C"))
		Expect(DecipherThis("65 119esi 111dl 111lw 108dvei 105n 97n 111ka")).Should(Equal("A wise old owl lived in an oak"))
		Expect(DecipherThis("84eh 109ero 104e 115wa 116eh 108sse 104e 115eokp")).Should(Equal("The more he saw the less he spoke"))
		Expect(DecipherThis("84eh 108sse 104e 115eokp 116eh 109ero 104e 104dare")).Should(Equal("The less he spoke the more he heard"))
		Expect(DecipherThis("87yh 99na 119e 110to 97ll 98e 108eki 116tah 119esi 111dl 98dri")).Should(Equal("Why can we not all be like that wise old bird"))
		Expect(DecipherThis("84kanh 121uo 80roti 102ro 97ll 121ruo 104ple")).Should(Equal("Thank you Piotr for all your help"))
	})
	It("Random Tests", func() {
		var randomInput string
		for i := 0; i < 100; i++ {
			randomInput = randomText()
			Expect(DecipherThis(randomInput)).Should(Equal(reference(randomInput)))
		}
	})
})

// Reference solution

func reference(text string) string {
	words := strings.Split(text, " ")
	result := []string{}
	for _, w := range words {
		nw := ""
		ll := ""
		lenOfNumber := 0
		for i, letter := range w {
			if letter >= 48 && letter <= 57 {
				ll += string(letter)
				lenOfNumber++
				if lenOfNumber == len(w) {
					nn, _ := strconv.Atoi(ll)
					nw += string(nn)
					ll = ""
				}
			} else {
				if ll != "" {
					nn, _ := strconv.Atoi(ll)
					nw += string(nn)
					ll = ""
				}
				if i == lenOfNumber {
					nw += string(w[len(w)-1])
				} else if i == len(w)-1 {
					nw += string(w[lenOfNumber])
				} else {
					nw += string(w[i])
				}
			}
		}
		result = append(result, nw)
	}
	return strings.Join(result, " ")
}

// Random input generator
func randomText() string {
	var words []string
	for i := 0; i < rand.Intn(20); i++ {
		words = append(words, randStringRunes(1+rand.Intn(5)))
	}
	return strings.Join(words, " ")
}

// Random string generator source: https://stackoverflow.com/a/31832326

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
var length = len(letterRunes)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(length)]
	}
	return string(b)
}
