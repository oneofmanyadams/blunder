package blunders

import (
	"time"
	"fmt"
)

// Blunder is the expanded version of an error.
//  - "Code" makes it easier to classify blunders. Should be a 1 word string.
//  - "Fatal" is used to determine if the blunder should halt the program.
//  - "Message" is the description provided when the specific blunder was reported.
//  - "Time" is used to help trace when the blunder happened.
type Blunder struct {
	Code string
	Message string
	Fatal bool
	Time time.Time
}

// NewBlunder generates a new Blunder type.
// Essentially mimics errors.New(string)
func NewBlunder(code string, message string, fatal bool, b_time time.Time) (blunder Blunder) {
	blunder.Code = code
	blunder.Message = message
	blunder.Fatal = fatal
	blunder.Time = b_time
	return
}

// Error allows a Blunder to be passed to any function that expects an error type.
// Does some basic formatting to make sure all Blunder fields are captured in the error string.
func (b *Blunder) Error() (error_string string) {
	var fatal_string string
	if b.Fatal {
		fatal_string = "FATAL"
	} else {
		fatal_string = "NON-FATAL"
	}
	error_string = fmt.Sprintf("%s, %s, \"%s\", %s", fatal_string, b.Code, b.Message, b.Time.Format("2006-01-02 15:04:05"))
	return
}