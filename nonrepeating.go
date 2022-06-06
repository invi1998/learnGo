package main

import "fmt"

// 计算最长非重复字符子串
func lengthOfNonRepeatingSubStr(s string) int {
	lastOccurred := make(map[byte]int)
	start := 0
	maxLength := 0

	for i, ch := range []byte(s) {

		if lastI, ok := lastOccurred[ch]; ok && lastI >= start {
			start = lastOccurred[ch] + 1
		}

		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}

		lastOccurred[ch] = i
	}

	return maxLength
}

func main() {
	fmt.Println(
		lengthOfNonRepeatingSubStr("abcabcbb"),           // 3
		lengthOfNonRepeatingSubStr("asdfdagsdwdgsouios"), // 7
		lengthOfNonRepeatingSubStr("pwwkew"),             // 3
	)
}
