package template

import "testing"

func TestHello(t *testing.T) {
	if Hello() != "Hello!" {
		t.Fatal("Didn't return the right string.")
	}
}

func TestAddRomanIs(t *testing.T) {
	added, _ := RomanNumeral("I").Add("I")
	if added != RomanNumeral("II") {
		t.Fatal("I+I should == II")
	}

	added3, _ := RomanNumeral("I").Add("II")
	if added3 != "III" {
		t.Fatal("I + II should == III")
	}
}

func TestAddRomanV(t *testing.T) {
	added4, _ := RomanNumeral("I").Add("III")
	if added4 != "IV" {
		t.Fatal("I+III should == IV")
	}

	added5, _ := RomanNumeral("III").Add("III")
	if added5 != "VI" {
		t.Fatal("III+III should == VI")
	}

	added6, _ := RomanNumeral("IV").Add("I")
	if added6 != "V" {
		t.Fatal("IV+I should == V")
	}

	added7, _ := RomanNumeral("V").Add("I")
	if added7 != "VI" {
		t.Fatal("V+I should == VI")
	}
}

func TestAddRomanIVPlusVI(t *testing.T) {
	added8, _ := RomanNumeral("IV").Add("VI")
	if added8 != "X" {
		t.Fatal("IV+VI should == X")
	}
}

func TestFailIfNotRomanNumeral(t *testing.T) {
	_, err := RomanNumeral("Q").Add("I")
	if err == nil {
		t.Fatal("Q + I should throw error")
	}
}


