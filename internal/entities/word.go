package entities

type rank struct {
	words map[string]int
	lang  string
}

func (r rank) getSentenceFreq(s []string) []int {
	var ranks []int
	for _, word := range s {
		ranks = append(ranks, r.words[word])
	}
	return ranks
}
