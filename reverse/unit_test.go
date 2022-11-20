package main

import "testing"

func Test(t *testing.T) {

	testCases := make(map[string]string)

	testCases["example00.txt"] = "Hello World"
	testCases["example01.txt"] = "123"
	testCases["example02.txt"] = "#=\\["
	testCases["example03.txt"] = "something&234"
	testCases["example04.txt"] = "abcdefghijklmnopqrstuvwxyz"
	testCases["example05.txt"] = "\\!\" #$%&'()*+,-./"
	testCases["example06.txt"] = ":;{=}?@"
	testCases["example07.txt"] = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	testCases["exampleLowerUpper.txt"] = "LowErUpPer"
	testCases["exampleLowerNumbersSpaces.txt"] = "1low er 2upper"
	testCases["exampleSpecial.txt"] = "spesi@l {()} ch*"
	testCases["exampleLowerUpperSpacesNumbers.txt"] = "2Lower 1Upper"

	for key, value := range testCases {
		path := key
		want := value
		result := convertText(path)
		if want != result {
			t.Fatalf(`convertText(path) fail.
Want: %v
Found: %v`, want, result)
		}
	}

}
