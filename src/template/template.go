package template

//Hello says hello
func Hello() string {
	return "Hello!"
}

type RomanNumeral string

// Adds roman numerals together (Ex. I + III = IV)
func AddRoman(s RomanNumeral, t RomanNumeral) RomanNumeral {
	IsByLiteral := map[RomanNumeral]RomanNumeral{
		"I":"I",
		"II":"II",
		"III":"III",
		"IV":"IIII",
		"V":"IIIII",
		"VI":"IIIIII",
	}
	LiteralsByI := map[RomanNumeral]RomanNumeral{
		"IIIIII":"VI",
		"IIIII":"V",
		"IIII":"IV",
		"III":"III",
		"II":"II",
		"I":"I",
	}
    Is := IsByLiteral[s] + IsByLiteral[t]
	return LiteralsByI[Is]
}