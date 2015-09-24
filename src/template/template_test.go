package template

import "testing"

func TestHello(t *testing.T) {
	if Hello() != "Hello!" {
		t.Fatal("Didn't return the right string.")
	}
}

func TestAddRomanIs(t *testing.T) {
	if RomanNumeral("I").Add("I") != "II" {
		t.Fatal("I+I should == II")
	}

	if RomanNumeral("I").Add("II") != "III" {
		t.Fatal("I + II should == III")
	}
}

func TestAddRomanV(t *testing.T) {
	if RomanNumeral("I").Add("III") != "IV" {
		t.Fatal("I+III should == IV")
	}

	if RomanNumeral("III").Add("III") != "VI" {
		t.Fatal("III+III should == VI")
	}

	if RomanNumeral("IV").Add("I") != "V" {
		t.Fatal("IV+I should == V")
	}

	if RomanNumeral("V").Add("I") != "VI" {
		t.Fatal("V+I should == VI")
	}
}

//func TestAddRomanIVPlusVI(t *testing.T) {
//	if RomanNumeral("IV").Add("VI") != "X" {
//		t.Fatal("IV+VI should == X")
//	}
//}


