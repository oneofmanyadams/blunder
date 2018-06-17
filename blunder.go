package blunders

import (
	"fmt"
)

// Blunder is the expanded version of an error.
//  - "code" makes it easier to classify blunders.
//  - "codeName" provides a more human friendly way of looking at a "code".
//  - "fatal" is used to determine if the blunder should halt the program.
//  - "message" is the description provided when the specific blunder was reported.
type Blunder struct {
	Code int
	CodeName string
	Fatal bool
	Message string
}

// NewBlunder generates a new Blunder type.
// Essentially mimics errors.New(string)
func NewBlunder(code int, code_name string, fatal bool, message string) (b Blunder) {
	b.Code = code
	b.CodeName = code_name
	b.Fatal = fatal
	b.Message = message
	return
} 

// Error allows a Blunder to be passed to any function that expects an error type.
// Does some basic formatting to make sure all Blunder fields are captured in the error string.
func (b *Blunder) Error() (error_string string) {
	var fatal_string string
	if b.Fatal {
		fatal_string = "FATAL BLUNDER"
	} else {
		fatal_string = "NON-FATAL BLUNDER"
	}

	error_string = fmt.Sprintf("%s enountered, CODE: %d (%s), \"%s\"", fatal_string, b.Code, b.CodeName, b.Message)
	return
}