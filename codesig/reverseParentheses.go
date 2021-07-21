import . "strings"
func pop(array []string) string {
	length := len(stack) - 1
	value := array[length]
	array = array[:length]
	return value
}

func reverseInParentheses(inputString string) {
	if len(inputString) < 3 {
		return ""
	}

	stack := make([]string, 0, 20)

	for _, character := inputString {
		if string(character) == ")" {
			temporary := ""

			value := pop(stack)

			for {
				if value == "(" {
					break
				}

				temporary += value
				value = pop(stack)

				if len(stack) == 0 {
					break
				}
			}

			if value == "(" {
				arrayT := Split(temporary)

				for _, r := range arrayT {
					stack = append(stack, string(r))
				}
				temporary = ""
			}
		} else {
			stack = append(stack, character)
		}
	}
	return Join(stack)
}