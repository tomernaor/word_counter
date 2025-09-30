package analyzer

import "strings"

type lexicon struct {
	valid   map[string]struct{}
	inValid map[string]struct{}
}

// TODO: will it be better to set size for maps?
func emptyLexicon() *lexicon {
	return &lexicon{
		valid:   make(map[string]struct{}),
		inValid: make(map[string]struct{}),
	}
}

func newLexiconFromFile(filename string) (*lexicon, error) {
	lexicon := emptyLexicon()

	err := readByLine(filename, func(word string) error {
		lexicon.add(word)

		return nil
	})

	if err != nil {
		return nil, err
	}

	return lexicon, nil
}

func (d *lexicon) add(word string) {
	word = normalize(word)

	if isValid(word) {
		d.valid[word] = struct{}{}
	} else {
		d.inValid[word] = struct{}{}
	}
}

func (d *lexicon) exists(word string) bool {
	_, exists := d.valid[word]
	return exists
}

func (d *lexicon) totalValid() int {
	return len(d.valid)
}

func (d *lexicon) totalInValid() int {
	return len(d.inValid)
}

func normalize(word string) string {
	return strings.ToLower(word)
}

func isValid(word string) bool {
	if len(word) < 3 {
		return false
	}

	for _, ch := range word {
		if (ch < 'A' || ch > 'Z') && (ch < 'a' || ch > 'z') {
			return false
		}
	}

	return true
}
