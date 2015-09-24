package template
import "errors"

//Hello says hello
func Hello() string {
	return "Hello!"
}

type RomanNumeral string

func (r RomanNumeral) Valid() bool {
	_, ok := isByLiteral[r]
	return ok
}

// Adds roman numerals together (Ex. I + III = IV)
func (r RomanNumeral) Add(s RomanNumeral) (RomanNumeral, error) {
	if s.Valid() && r.Valid() {
		is := isByLiteral[r] + isByLiteral[s]
		return literalsByI[is], nil
	}
	return "", errors.New("One of these is not a roman numeral")
}

var isByLiteral = map[RomanNumeral]RomanNumeral{
	"I":   "I",
	"II":  "II",
	"III": "III",
	"IV":  "IIII",
	"V":   "IIIII",
	"VI":  "IIIIII",
	"X":   "IIIIIIIIII",
}

var literalsByI = map[RomanNumeral]RomanNumeral{
	"IIIIIIIIII": "X",
	"IIIIII":     "VI",
	"IIIII":      "V",
	"IIII":       "IV",
	"III":        "III",
	"II":         "II",
	"I":          "I",
}
