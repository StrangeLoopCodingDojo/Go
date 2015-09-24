package template

//Hello says hello
func Hello() string {
	return "Hello!"
}

type RomanNumeral string

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

// Adds roman numerals together (Ex. I + III = IV)
func AddRoman(s RomanNumeral, t RomanNumeral) RomanNumeral {
	Is := isByLiteral[s] + isByLiteral[t]
	return literalsByI[Is]
}
