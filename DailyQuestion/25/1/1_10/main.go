package __10

func contains(arr1, arr2 []int) bool {
	for i := 0; i < len(arr2); i++ {
		if arr1[i] < arr2[i] {
			return false
		}
	}
	return true
}

func validSubstringCount(word1 string, word2 string) (result int64) {
	chars1, chars2 := make([]int, 26), make([]int, 26)
	for _, c := range word2 {
		chars2[c-'a']++
	}

	for i, j := int64(0), int64(0); j < int64(len(word1)); j++ {
		{
			c := word1[j]
			chars1[c-'a']++
		}
		for contains(chars1, chars2) {
			result += int64(len(word1)) - j

			c := word1[i]
			chars1[c-'a']--
			i++
		}
	}

	return result
}
