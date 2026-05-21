func isAnagram(s string, t string) bool {
	if len(s) != len(t){
		return false
	}
	loweredS := strings.ToLower(s)
	loweredT := strings.ToLower(t)

	mapS := make(map[rune]int)
	mapT := make(map[rune]int)

	for _, value := range loweredS{
		mapS[value]++
	}

	for _, value := range loweredT{
		mapT[value]++
	}

	fmt.Printf("MAP S: %v\n", mapS)
	fmt.Printf("MAP T: %v\n", mapT)

	for key, value := range mapS{
		if mapT[key] != value {
        return false
    }
	}
	return true
}