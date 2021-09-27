package string_modifer

import (
	"testing"
)


func TestAlterNateUpperCaseConverter_ConvertToUpperCase(t *testing.T) {
	var alternateUppercaseConverter IAlternateUpperCaseConverter
	alternateUppercaseConverter= NewAlternateUpperCaseConverter()
	input := "hello world"
	expectedResult := "hElLo wOrLd" 
	result := string(alternateUppercaseConverter.ConvertToAlternateUpperCase([]byte(input)))
	if result !=expectedResult {
		t.Errorf("Expected alternate uppercase to be %s but got %s\n", result, expectedResult)
	}
}