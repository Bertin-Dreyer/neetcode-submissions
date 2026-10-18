func isValid(s string) bool {
    bracket := map[rune]rune{'(':')' , '{':'}' , '[':']'}

	stack := []rune{}
	sRunes := []rune(s)

	for _, char := range sRunes {
		// if we have seen the opening bracket,
		// we add the closing tag into the stack so we can
		// pop it later in order.
		if _, seen := bracket[char] ; seen{
			stack = append(stack, bracket[char])
		} else {
			// check if the stack can pop on a closing bracket

			//if the length of the stack is 0 then its false.
			// OR
			// if the closing bracket isnt the bracket we have at the top of the stack:
			if len(stack) == 0 || char != stack[len(stack) - 1]  {
				return false
			}
			// pop the bracket from the stack
			stack = stack[:len(stack) - 1]
		}
	}
	// bool eval to see if the stack is cleared and good.
	return len(stack) == 0
}
