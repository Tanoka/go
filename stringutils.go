package stringutils

var spaces [4]rune

func init() {
	spaces[0] = ' '
	spaces[1] = '\n'
	spaces[2] = '\t'
	spaces[3] = '\v'
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

// 804 ns/op vs strings.Fields 184 ns/op
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

// 242 ns/op vs strings.Fields 184 ns/op
func SplitIntoWordsFast(s string) (res []string) {
	var ini int = -1
	for i, c := range s {
		if !isSpace(c) && ini == -1 {
			ini = i
		} else if isSpace(c) && ini > 0 {
			res = append(res, s[ini:i])
			ini = -1
		}
	}
	if ini > 0 { //Last word if string not ending with space or new line
		res = append(res, s[ini:])
	}
	return
}
