package stringutils

var spaces [2]rune

func init() {
	spaces[0] = ' '
	spaces[1] = '\n'
}

func isSpace(c rune) bool {
	for _, s := range spaces {
		if c == s {
			return true
		}
	}
	return false
}

func CountChars(s string) int {
	r := []rune(s)
	return len(r)
}

func SplitIntoWords(s string) (res []string) {
	var word []rune
	for _, c := range s {
		if !isSpace(c) {
			word = append(word, c)
		} else {
			if len(word) > 0 {
				res = append(res, string(word))
				word = nil
			}
		}
	}
	if len(word) > 0 { //Last word if string not ending with space or new line
		res = append(res, string(word))
	}
	return
}
