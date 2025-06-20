package utils

func FindStringIndex(text string, subStr string) [][]int {
	var indexes [][]int

	for i := range len(text) {
		if text[i] == subStr[0] {
			end := len(subStr) + i
			if text[i:end] == subStr {
				indexes = append(indexes, []int{i, end})
			}
		}
	}

	return indexes

}
