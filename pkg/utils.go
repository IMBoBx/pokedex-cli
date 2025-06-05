package pkg

import (
	"strings"
)

const (
	Base       string = "https://pokeapi.co/api/v2/"
	LatestGame string = "scarlet-violet"
)

func ToTitle(s string) string {

	words := strings.Split(strings.ReplaceAll(s, " ", "-"), "-")
	if len(words) == 1 {
		if w := words[0]; len(w) > 2 {
			return strings.ToUpper(w[0:1]) + strings.ToLower(w[1:])
		} else {
			return strings.ToUpper(w)
		}
	}

	for i, w := range words {
		words[i] = ToTitle(w)
	}

	return strings.Join(words, " ")
}
