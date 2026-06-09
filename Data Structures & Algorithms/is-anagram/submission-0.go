func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	
	wordMap := make(map[string]int)

	for _, char := range s {
		wordMap[string(char)]++
	}
	for _, char := range t {
		wordMap[string(char)]--
	}

	for _, value := range wordMap {
        if value != 0 {
            return false 
        }
    }

	return true
}
