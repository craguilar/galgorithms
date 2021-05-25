package mstrings

import "strings"

func TwoChars(n int, text string) int {

	//
	var count [26]int
	max := 0

	for i := 0; i < n; i++ {
		count[(text[i]-'a')]++
	}
	for i := 0; i < 26; i++ {
		if count[i] > 0 {
			for j := i + 1; j < 26; j++ {
				if count[j] > 0 {
					s := keepChars(text, i+'a', j+'a')
					if checkAlternate(s) && len(s) > max {
						max = len(s)
					}
				}

			}
		}

	}
	return n
}

func keepChars(s string, a int, b int) string {
	f := ""
	for i := 0; i < len(s); i++ {
		if int(s[i]) == a || int(s[i]) == b {
			f = f + string(s[i])
		}
	}
	return f
}

func checkAlternate(s string) bool {
	isAlternate := true
	first := rune(s[0])
	for i := 1; i < len(s); i++ {
		if first == rune(s[i]) {
			return false
		}
		first = rune(s[i])
	}
	return isAlternate
}

func CamelCase(text string) int {
	// 0 length string
	if text == "" {
		return 0
	}
	c := 1
	//iterate over the string
	for i, _ := range text {
		if string(text[i]) == strings.ToUpper(string(text[i])) {
			c++
		}
	}
	return c
}
