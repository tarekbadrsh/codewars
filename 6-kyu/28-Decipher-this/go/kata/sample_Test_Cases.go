package kata_test

import (
	. "codewarrior/kata"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Test Example", func() {
	It("Sample Tests", func() {
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
})
