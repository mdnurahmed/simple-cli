package string_writer

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/mock"
)


type MockCSVWriter struct {
	mock.Mock
}

func (m *MockCSVWriter) WriteToCSV(s []byte) error {
	args := m.Called(s)
	return args.Error(0)
}

func TestCSVWriter_WriteToCSV_ErrorCase(t *testing.T){
	csvWriter := MockCSVWriter{}
	expectedError := fmt.Errorf("An error") 
	csvWriter.On("WriteToCSV",[]byte("Throw An Error")).Return(expectedError)

	input := "Throw An Error"
	err := csvWriter.WriteToCSV([]byte(input))
	if err != expectedError {
		t.Errorf("Expected error - %s but got error - %s\n", expectedError.Error(),err.Error())
	}
}	


func TestCSVWriter_WriteToCSV_HappyCase(t *testing.T){
	csvWriter := MockCSVWriter{}
	csvWriter.On("WriteToCSV",[]byte("hello world")).Return(nil)
	input := "hello world"
	err := csvWriter.WriteToCSV([]byte(input))
	if err != nil {
		t.Errorf("Expected no error but got error - %s\n",err.Error())
	}

}	