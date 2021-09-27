package main

import (
	"bufio"
	. "cli/string_modifier"
	. "cli/string_writer"
	"fmt"
	"os"
)

func main() {
	var uppercaseConverter IUpperCaseConverter
	var alternateUppercaseConverter IAlternateUpperCaseConverter
	var csvWriter ICSVWriter
	var input string
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("\nEnter a input string : ")
		scanner.Scan() 
		input = scanner.Text()
		if len(input) == 0 {
			fmt.Println("Input length cannot be zero")
			continue
		}
		uppercaseConverter = NewUpperCaseConverter()
		fmt.Println(string(uppercaseConverter.ConvertToUpperCase([]byte(input))))

		alternateUppercaseConverter = NewAlternateUpperCaseConverter()
		fmt.Println(string(alternateUppercaseConverter.ConvertToAlternateUpperCase([]byte(input))))

		csvWriter = NewCSVWriter()
		err := csvWriter.WriteToCSV([]byte(input))
		if err != nil {
			fmt.Printf("Error occured when writing to csv - %s \n", err.Error())
		} else {
			fmt.Println("CSV created!")
		}
	}

}
