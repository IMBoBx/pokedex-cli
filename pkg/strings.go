package pkg

import (
	"strings"
)

const Base string = "https://pokeapi.co/api/v2/"


func ToTitle(s string) string {
	return strings.ToUpper(s[0:1]) + s[1:]
}