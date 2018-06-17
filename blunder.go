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
	code int
	codeName string
	fatal bool
	message string
}

// NewBlunder generates a new Blunder type.
// Essentially mimics errors.New(string)
func NewBlunder(code int, code_name string, fatal bool, message string) (b Blunder) {
	b.code = code
	b.codeName = code_name
	b.fatal = fatal
	b.message = message
	return
} 

// Error allows a Blunder to be passed to any function that expects an error type.
// Does some basic formatting to make sure all Blunder fields are captured in the error string.
func (b *Blunder) Error() (error_string string) {
	var fatal_string string
	if b.fatal {
		fatal_string = "FATAL BLUNDER"
	} else {
		fatal_string = "NON-FATAL BLUNDER"
	}

	error_string = fmt.Sprintf("%s enountered, CODE: %d (%s), \"%s\"", fatal_string, b.code, b.codeName, b.message)
	return
}