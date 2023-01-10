package main

import (
	"strings"
)

type sent struct {
	lang string
	sent string
}

type pair struct {
	sent_a sent
	sent_b sent
}

func toSentStruct(rawPair []string) pair {
	raw_a := rawPair[1]
	raw_b := rawPair[2]

	lang_a := strings.Split(raw_a, "\"")[1]
	lang_b := strings.Split(raw_b, "\"")[1]

	sentence_a := strings.Split(strings.Split(raw_a, "<seg>")[1], "</seg>")[0]
	sentence_b := strings.Split(strings.Split(raw_b, "<seg>")[1], "</seg>")[0]

	return pair{
		sent{
			lang_a,
			sentence_a,
		},
		sent{
			lang_b,
			sentence_b,
		},
	}
}

func getStructList(sentences []string) []pair {
	pairs := make([]pair, 0)
	for i := 0; i < len(sentences)-2; i = i + 4 {
		pairs = append(pairs, toSentStruct(sentences[i:i+4]))
	}
	return pairs
}
