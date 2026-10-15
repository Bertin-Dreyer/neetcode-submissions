func isValid(s string) bool {
    stack := []rune{}
	

	check := map[rune]rune{
		'{':'}',
		'[':']',
		'(':')',
	}

	for _, bracket := range s {

		if _, ok := check[bracket]; ok{
			stack = append(stack,bracket)
		}else {
			if len(stack) == 0 {
				return false
			}

			top := stack[len(stack) -1]
			want := check[top]

			if want != bracket {
				return false
			}

			stack = stack[:len(stack) -1]
		}
	}
	return len(stack) == 0
}