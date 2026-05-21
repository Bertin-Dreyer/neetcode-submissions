func isAnagram(s string, t string) bool {
	if len(s) != len(t){
		return false
	}

	s = strings.ToLower(s)
    t = strings.ToLower(t)

	mapS := make(map[rune]int, len(s))
	mapT := make(map[rune]int, len(t))

	for _, value := range s{
		mapS[value]++
	}

	for _, value := range t{
		mapT[value]++
	}

	for key, value := range mapS{
		if mapT[key] != value {
        return false
    	}
	}
	return true
}