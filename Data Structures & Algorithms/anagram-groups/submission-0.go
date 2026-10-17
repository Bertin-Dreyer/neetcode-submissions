import "slices"
func groupAnagrams(strs []string) [][]string {

	outputList := [][]string{}
	strMap := make(map[string][]string)

	for _, str := range strs {
		key := sortStr(str)
		strMap[key] = append(strMap[key], str)
	}

	for _ , slice := range strMap {
		outputList = append(outputList, slice)
	}

	return outputList

}

func sortStr(s string) string {
	sRunes := []rune(s)
	slices.Sort(sRunes)

	return string(sRunes)
}