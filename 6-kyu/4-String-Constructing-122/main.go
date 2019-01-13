package main

func main() {

	StringConstructing("aba", "abaa")
	// "" -> "aba" -> "abaaba" -> "abaab" -> "abaa"

	StringConstructing("a", "aaa")
	// "" -> "a" -> "aa" -> "aaa"

	StringConstructing("aba", "abbabba")
	//  "" -> "aba" -> "ab" -> "ababa" -> "abba" -> "abbaaba" -> "abbaab".

	StringConstructing("aba", "a")
	// "" -> "aba" -> "ab" -> "a".

	StringConstructing("a", "a")
	// "" -> "a"

	StringConstructing("abcdefgh", "hgfedcba")
	StringConstructingBestPractices("abcdefgh", "hgfedcba")
}

// StringConstructing :
func StringConstructing(a, s string) int {
	final := ""
	op := 0
	for final != s {
		if len(s) > len(final) {
			final = final + a
			op++
		}
		k := 0
		for len(final) > k {
			if len(s) <= k && len(final) >= k {
				op += len(final[k:])
				final = final[:k]
				break
			}
			if final[k] != s[k] && len(final) > k {
				final = final[:k] + final[k+1:]
				op++
			} else {
				k++
			}
		}
	}
	return op
}

// StringConstructingBestPractices :
func StringConstructingBestPractices(a, s string) int {
	var i, ops int
	var curr string

	for curr != s {
		for i < len(curr) && i < len(s) && s[i] == curr[i] {
			i++
		}

		if i >= len(curr) {
			curr += a
		} else {
			curr = curr[:i] + curr[i+1:len(curr)]
		}

		ops++
	}

	return ops
}
