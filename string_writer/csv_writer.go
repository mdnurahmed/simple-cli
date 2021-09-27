package string_writer

import (
	"encoding/csv"
	"os"
)

//ICSVWriter is the interface for writing a string to csv . This interface can have many implementation.
type ICSVWriter interface {
	WriteToCSV(s []byte) error
}

//CSVWriter implements ICSVWriter
type CSVWriter struct{}

func (c *CSVWriter) WriteToCSV(s []byte) error {
	csvFile, err := os.Create("csv.csv")
	if err != nil {
		return err
	}
	defer csvFile.Close()
	csvwriter := csv.NewWriter(csvFile)
	defer csvwriter.Flush()
	data  := make([][]string,1)
	for _, letter := range s {
		data[0]  = append(data[0], string(letter))
	}
	for _, value := range data {
        err := csvwriter.Write(value)
        if err != nil {
			return err 
		}
    }
	return nil
}

func NewCSVWriter() *CSVWriter {
	return &CSVWriter{}
}
