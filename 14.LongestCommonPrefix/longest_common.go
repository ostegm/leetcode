package main

import "fmt"

func comparePrefixes(s1, s2 string) string {
	var totalChars int
	if len(s1) > len(s2) {
		totalChars = len(s2)
	} else {
		totalChars = len(s1)
	}

	if totalChars == 0 {
		return ""
	}

	for i := 0; i < totalChars+1; i++ {
		if s1[:i] != s2[:i] {
			return s1[:i-1]
		}
	}
	return s1[:totalChars]

}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	longest := strs[0]
	for i := 0; i < len(strs); i++ {
		longest = comparePrefixes(longest, strs[i])
		if longest == "" {
			return longest
		}
	}
	return longest
}

func main() {
	strs := []string{"dog", "doggie", "dogggggg"}
	got := longestCommonPrefix(strs)
	fmt.Println(got)
}
