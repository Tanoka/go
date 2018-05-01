package stringutils

import "testing"


func TestCountChars(t *testing.T) {
	var flagtests = []struct {
		st string
		cu int
	}{
		{"abc def", 7},
		{"ábc déf", 7},
		{"ab\n\td", 5},
		{"ab \u21D4", 4},
	}


	for _, te := range flagtests {
		if re := CountChars(te.st); re != te.cu {
			t.Errorf("No equals. Parameter:%q,  Expected: %d, actual: %d", te.st, te.cu, re)
		}

	}

}

func TestSplitIntoWords(t *testing.T) {
	type words []string
	var flagtests = []struct {
		st string
		re words
	}{
		{"one", words{"one"}},
		{"uno dís",words{"uno", "dís"}},
		{"    ",words{}},
		{"  unó    ",words{"unó"}},
		{"uno   dós ",words{"uno","dós"}},
	}

	for _, te := range flagtests {
		re := SplitIntoWords(te.st)
		for i, st := range re {
			if st != te.re[i] {
				t.Errorf("Not equals. Expected:%v, Actual:%v", te.re, re)
				break
			}

		}
	}
}
