func lengthOfLongestSubstring(s string) int {
    start := 0
    maxLength := 0
    charMap := make(map[rune]int)

    for end, char := range s {
        if index, ok := charMap[char]; ok && index >= start {
            start = index + 1
        }
        charMap[char] = end
        currentLength := end - start + 1
        if currentLength > maxLength {
            maxLength = currentLength
        }
    }
    return maxLength
}
