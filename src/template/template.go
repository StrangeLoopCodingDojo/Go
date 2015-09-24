package template

//Hello says hello
func Hello() string {
	return "Hello!"
}

type RomanNumeral string

// Adds roman numerals together (Ex. I + III = IV)
func (r RomanNumeral) Add(s RomanNumeral) RomanNumeral {
	Is := isByLiteral[r] + isByLiteral[s]
	return literalsByI[Is]
}

var isByLiteral = map[RomanNumeral]RomanNumeral{
	"I":   "I",
	"II":  "II",
	"III": "III",
	"IV":  "IIII",
	"V":   "IIIII",
	"VI":  "IIIIII",
}

var literalsByI = map[RomanNumeral]RomanNumeral{
	"IIIIII": "VI",
	"IIIII":  "V",
	"IIII":   "IV",
	"III":    "III",
	"II":     "II",
	"I":      "I",
}
