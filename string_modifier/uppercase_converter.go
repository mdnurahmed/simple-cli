package string_modifer

import "unicode"

//IUpperCaseConverter is the interface for the uppercase converter. This interface can have many implimentation.
type IUpperCaseConverter interface {
	ConvertToUpperCase(s []byte) []byte
}

// UpperCaseConverter impliments IUpperCaseConverter
type UpperCaseConverter struct{}

func (u *UpperCaseConverter) ConvertToUpperCase(s []byte) []byte {
	for i := 0; i < len(s); i++ {
		s[i] = byte(unicode.ToUpper(rune(s[i])))
	}
	return s
}

func NewUpperCaseConverter() *UpperCaseConverter {
	return &UpperCaseConverter{}
}
