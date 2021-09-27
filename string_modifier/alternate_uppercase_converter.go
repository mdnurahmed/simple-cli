package string_modifer

import "unicode"

//IAlternateUpperCaseConverter is the interface for the alternate uppercase converter. This interface can have many implementation.
type IAlternateUpperCaseConverter interface {
	ConvertToAlternateUpperCase(s []byte) []byte
}

// AlternateUpperCaseConverter implements IAlternateUpperCaseConverter
type AlternateUpperCaseConverter struct{}

func (u *AlternateUpperCaseConverter) ConvertToAlternateUpperCase(s []byte) []byte {
	for i := 1; i < len(s); i += 2 {
		s[i] = byte(unicode.ToUpper(rune(s[i])))
	}
	return s
}

func NewAlternateUpperCaseConverter() *AlternateUpperCaseConverter {
	return &AlternateUpperCaseConverter{}
}
