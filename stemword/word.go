/*
	This package implements various string-like methods
	over slices of runes.  It is similar to the
	exp/utf8string package.
*/
package stemword

// Word represents a word that is going to be stemmed.
// 
type Word struct {

	// A slice of runes
	RS []rune

	// The index in RS where the R1 region begins
	R1i int

	// The index in RS where the R2 region begins
	R2i int
}

// Create a new Word struct
func New(in string) (word *Word) {
	word = &Word{RS: []rune(in)}
	word.R1i = len(word.RS)
	word.R2i = len(word.RS)
	return
}

// Resets R1i and R2i to ensure they 
// are within bounds of the current rune slice.
func (w *Word) resetR1R2() {
	rsLen := len(w.RS)
	if w.R1i > rsLen {
		w.R1i = rsLen
	}
	if w.R2i > rsLen {
		w.R2i = rsLen
	}
}

// Return a slice of w.RS, allowing the start
// and stop to be out of bounds.
//
func (w *Word) slice(start, stop int) []rune {
	startMin := 0
	if start < startMin {
		start = startMin
	}
	max := len(w.RS) - 1
	if start > max {
		start = max
	}
	if stop > max {
		stop = max
	}
	return w.RS[start:stop]
}

// Return the R1 region as a slice of runes
func (w *Word) R1() []rune {
	return w.RS[w.R1i:]
}

// Return the R1 region as a string
func (w *Word) R1String() string {
	return string(w.R1())
}

// Return the R2 region as a slice of runes
func (w *Word) R2() []rune {
	return w.RS[w.R2i:]
}

// Return the R2 region as a string
func (w *Word) R2String() string {
	return string(w.R2())
}

// Return the Word as a string
func (w *Word) String() string {
	return string(w.RS)
}

// Return the first prefix found or the empty string.
func (w *Word) FirstPrefix(prefixes ...string) string {
	found := false
	rsLen := len(w.RS)

	for _, prefix := range prefixes {
		found = true
		for i, r := range prefix {
			if i > rsLen-1 || (w.RS)[i] != r {
				found = false
				break
			}
		}
		if found {
			return prefix
		}
	}
	return ""
}

// Return the first suffix found or the empty string.
func (w *Word) FirstSuffix(sufficies ...string) (suffix string) {
	rsLen := len(w.RS)
	for _, suffix := range sufficies {
		numMatching := 0
		suffixRunes := []rune(suffix)
		suffixLen := len(suffixRunes)
		for i := 0; i < rsLen && i < suffixLen; i++ {
			if w.RS[rsLen-i-1] != suffixRunes[suffixLen-i-1] {
				break
			} else {
				numMatching += 1
			}
		}
		if numMatching == suffixLen {
			return suffix
		}
	}
	return ""
}