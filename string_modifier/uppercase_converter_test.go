package string_modifer

import "testing"

func TestUpperCaseConverter_ConvertToUpperCase(t *testing.T) {
	var uppercaseConverter IUpperCaseConverter
	uppercaseConverter = NewUpperCaseConverter()
	input := "hello world"
	expectedResult := "HELLO WORLD"
	result := string(uppercaseConverter.ConvertToUpperCase([]byte(input)))
	if result != expectedResult {
		t.Errorf("Expected alternate uppercase to be %s but got %s\n", result, expectedResult)
	}
}