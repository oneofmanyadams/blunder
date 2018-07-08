package blunders

import (
	"fmt"
	"time"
	"io"
)
//////////////////////////////////////////////////////////////////
// Blunder Code
//////////////////////////////////////////////////////////////////

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
		fatal_string = "FATAL BLUNDER"
	} else {
		fatal_string = "NON-FATAL BLUNDER"
	}
	error_string = fmt.Sprintf("%s, %s enountered, %s, \"%s\"", b.Time.String(), fatal_string, b.Code, b.Message)
	return
}

//////////////////////////////////////////////////////////////////
// BlunderBus Code
//////////////////////////////////////////////////////////////////

type BlunderBus struct {
	Blunders []Blunder
	HasFatal bool
}

func NewBlunderBus() (bb *BlunderBus) {
	var blunder_bus BlunderBus
	bb = &blunder_bus
	bb.HasFatal = false // this really isn't necessary, buuuuttt... ¯\_(ツ)_/¯
	return
}

//////////////////////////////////////////////////////////////////
// BlunderBus New Blunder Methods
//////////////////////////////////////////////////////////////////

func (bb *BlunderBus) New(code string, message string) {
	bb.newBase(code, message, false)
}

func (bb *BlunderBus) NewFatal(code string, message string) {
	bb.newBase(code, message, true)
}

func (bb *BlunderBus) newBase(code string, message string, fatal bool) {
	bb.Blunders = append(bb.Blunders, NewBlunder(code, message, fatal, time.Now()))
	if fatal {
		bb.HasFatal = true
	}
}

//////////////////////////////////////////////////////////////////
// BlunderBus Information Methods
//////////////////////////////////////////////////////////////////

func (bb BlunderBus) Fatals() (fatals []Blunder) {
	for _, blndr := range bb.Blunders {
		if blndr.Fatal {
			fatals = append(fatals, blndr)
		}
	}
	return
}

func (bb BlunderBus) NonFatals() (non_fatals []Blunder) {
	for _, blndr := range bb.Blunders {
		if !blndr.Fatal {
			non_fatals = append(non_fatals, blndr)
		}
	}
	return
}

func (bb BlunderBus) Codes(code string) (matching_codes []Blunder) {
	for _, blndr := range bb.Blunders {
		if blndr.Code == code {
			matching_codes = append(matching_codes, blndr)
		}
	}	
	return
}

func (bb BlunderBus) OrderedByCode() (blunder_groups []Blunder) {
	blunder_codes := bb.MappedByCode()
	for _, blunder_code_groups := range blunder_codes {
		for _, blndr := range blunder_code_groups {
			blunder_groups = append(blunder_groups, blndr)
		}
	}
	return
}

func (bb BlunderBus) MappedByCode() (blunder_groups map[string][]Blunder) {
	blunder_groups = make(map[string][]Blunder)
	for _, blndr := range bb.Blunders {
		blunder_groups[blndr.Code] = append(blunder_groups[blndr.Code], blndr)
	}
	return
}

//////////////////////////////////////////////////////////////////
// BlunderBus Utility Methods
//////////////////////////////////////////////////////////////////

func BlunderSliceAsString(blunder_slice []Blunder) (blunder_string string) {
	for _, blndr := range blunder_slice {
		blunder_string = blunder_string + blndr.Error() + "\n"
	}
	return
}

func (bb BlunderBus) LogTo(writer io.Writer) {
	all_blunders := []byte(BlunderSliceAsString(bb.Blunders))
	writer.Write(all_blunders)
}