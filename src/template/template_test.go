package template

import "testing"

func TestHello(t *testing.T) {
	if Hello() != "Hello!" {
		t.Fatal("Didn't return the right string.")
	}
}

func TestAddRomanIs(t *testing.T) {
	if AddRoman("I", "I") != "II" {
		t.Fatal("I+I should == II")
	}

	if AddRoman("I", "II") != "III" {
		t.Fatal("I + II should == III")
	}
}

func TestAddRomanV(t *testing.T) {
	if AddRoman("I", "III") != "IV" {
		t.Fatal("I+III should == IV")
	}

	if AddRoman("III", "III") != "VI" {
		t.Fatal("III+III should == VI")
	}

	if AddRoman("IV", "I") != "V" {
		t.Fatal("IV+I should == V")
	}

	if AddRoman("V", "I") != "VI" {
		t.Fatal("V+I should == VI")
	}
}

//func TestAddRomanIVPlusVI(t *testing.T) {
//	if AddRoman("IV", "VI") != "X" {
//		t.Fatal("IV+VI should == X")
//	}
//}


