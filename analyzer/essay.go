package analyzer

type essay struct {
	urls []string
}

func emptyEssay() *essay {
	return &essay{
		urls: make([]string, 0),
	}
}

func newEssayFromFile(filename string) (*essay, error) {
	essay := emptyEssay()

	err := readByLine(filename, func(url string) error {
		essay.add(url)

		return nil
	})

	if err != nil {
		return nil, err
	}

	return essay, nil
}

func (d *essay) add(url string) {
	d.urls = append(d.urls, url)
}

func (d *essay) total() int {
	return len(d.urls)
}
