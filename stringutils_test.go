package stringutils

import "testing"
import "strings"

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
		{"    ",words{""}},
		{"  unó    ",words{"unó"}},
		{"úno ",words{"úno"}},
		{"úno   dós ",words{"úno","dós"}},
		{"uno\ndós ",words{"uno","dós"}},
		{"uno\tdós\n",words{"uno","dós"}},
	}

	for _, te := range flagtests {
		re := SplitIntoWords(te.st)
		for i, st := range re {
			if st != te.re[i] {
				t.Errorf("Not equals. Param:%q  Expected:%v, Actual:%v",te.st, te.re, re)
				break
			}

		}
	}
}

func TestSplitIntoWordsFast(t *testing.T) {
	type words []string
	var flagtests = []struct {
		st string
		re words
	}{
		{"one", words{"one"}},
		{"uno dís",words{"uno", "dís"}},
		{"    ",words{""}},
		{"  unó    ",words{"unó"}},
		{"úno ",words{"úno"}},
		{"uno   dós ",words{"uno","dós"}},
		{"uno\ndós ",words{"uno","dós"}},
		{"uno\tdós\n",words{"uno","dós"}},
	}

	for _, te := range flagtests {
		re := SplitIntoWordsFast(te.st)
		for i, st := range re {
			if st != te.re[i] {
				t.Errorf("Not equals. Param:%q  Expected:%v, Actual:%v",te.st, te.re, re)
				break
			}

		}
	}
}

func BenchmarkSplitIntoWords(b *testing.B) {
	for i:=0; i < b.N; i++ {
		SplitIntoWords("uno dós \ntres\tcuatro")
	}
}
func BenchmarkSplitIntoWordsFast(b *testing.B) {
	for i:=0; i < b.N; i++ {
		SplitIntoWordsFast("uno dós \ntres\tcuatro")
	}
}

func BenchmarkSplit(b *testing.B) {
	for i:=0; i < b.N; i++ {
		strings.Fields("uno dós \ntres\tcuatro")
	}
}


