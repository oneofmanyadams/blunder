// Package blunders provides more detailed and usable errors while still remaining simple to use.
package blunders

import (
	"fmt"
)

// Blunders is the main type for the blunders package.
//  - "Identifier" is name used to referance a speficic Blunders instance (such as when logging).
//  - "Codes" is list of key->value pairs to associate names to error codes.
//  - "Reported" is a list of every reported Blunder since the Blunders instance was created.
//  - "selfBlunders" is a list of any blunders encountered with the Blunders instance itself,
//  	such as trying to create multiple "Codes" that use the same id.
type Blunders struct {
	Identifier string
	Codes map[int]string
	Reported []Blunder
	selfBlunders []Blunder
}

// NewBlunders creates a new instance of Blunders.
// Sets the identifier name and initializes the "Codes" map.
// Creates Code[0] "Unregistered" default Code.
// Returns a new Blunders instance.
func NewBlunders(identifier string) (new_blunders Blunders) {
	new_blunders.Identifier = identifier
	new_blunders.Codes = make(map[int]string)
	new_blunders.RegisterCode(0, "UnregisteredBlunder")
	return
}

//////////////////////////////////////////////////////////////////
// Codes Functions
//////////////////////////////////////////////////////////////////

// RegisterCode creates a new blunder code and it's code name in "Codes".
// code_number "0" and code_name "UnregisteredBlunder" are reserved for the Blunders package. 
// It automatically checks to make sure the code or code name does not already exist.
// It will log a Blunder in "selfBlunders" if code_number or code_name already exists in "Codes".
// Returns true if the Code was created or false if it was not created.
func (b *Blunders) RegisterCode(code_number int, code_name string) (success bool) {
	success = true
	for existing_code_number, existing_code_name := range b.Codes {
		if existing_code_number == code_number {
			success = false
			b.newSelfBlunder(fmt.Sprintf("Attempted to use existing Code id \"%d\".", code_number))
			return
		}
		if existing_code_name == code_name {
			success = false
			b.newSelfBlunder(fmt.Sprintf("Attempted to use existing Code name \"%s\".", code_name))
			return
		}
	}

	b.Codes[code_number] = code_name
	return
}

// UnRegisterCode does not actually provide any functionality.
// I cannot think of any legitimate reason why one would need to dynamically un register a blunder Code.
// The only reason would be to create a new Code that uses the same id or name as an already existing Code,
// and having different Codes with the same id/name defeats the purpose of having ids/names int he first place.
// This is basically here just to remind me what a dumb idea this would be
// when I inevitably try to implement this in the future.
func (b *Blunders) UnRegisterCode(code_number int) {

}

//////////////////////////////////////////////////////////////////
// New Blunder Functions
//////////////////////////////////////////////////////////////////

// New is the standard function used to record a NON-FATAL blunder.
// This is designed for recording blunders that won't necissarily cause the program to crash.
// If a non-existing Code id is used, it will use the UnregisteredBlunder Code/CodeName
// and record a Blunder in selfBlunders.
func (b *Blunders) New(code int, message string) (blunder Blunder) {
	fatal := false
	b.newBlunderBase(code, fatal, message)
	return
}

// NewFatal is the standard function used to record a FATAL blunder.
// This is designed for recording blunders that will likely cause the program to crash.
// If a non-existing Code id is used, it will use the UnregisteredBlunder Code/CodeName
// and record a Blunder in selfBlunders.
func (b *Blunders) NewFatal(code int, message string) (blunder Blunder) {
	fatal := true
	b.newBlunderBase(code, fatal, message)
	return
}

// newBlunderBase the core function used to record a blunder.
// This provides a common base for both "FATAL" and "NON-FATAL" blunders.
// If a non-existing Code id is used, it will use the UnregisteredBlunder Code/CodeName
// and record a Blunder in selfBlunders.
func (b *Blunders) newBlunderBase(code int, fatal bool, message string) (blunder Blunder) {
	var code_id int
	var code_name string

	if name, exists := b.Codes[code]; exists {
		code_id = code
		code_name = name
	} else {
		code_id = 0
		code_name = b.Codes[0]
		b.newSelfBlunder(fmt.Sprintf("Attempted to use unregistered Code id \"%d\".", code))
	}

	blunder = NewBlunder(code_id, code_name, fatal, message)
	b.Reported = append(b.Reported, blunder)
	
	return
}

// newSelfBlunder is used to record any blunders encountered with the Blunders package itself.
// By default, it uses the blunder Code "0" and the CodeName "self_blunder".
// All self-Blunders are considered non-fatal.
func (b *Blunders) newSelfBlunder(message string) {
	b.selfBlunders = append(b.selfBlunders, NewBlunder(0, "self_blunder", false, message))
}

//////////////////////////////////////////////////////////////////
// Utility Functions
//////////////////////////////////////////////////////////////////
func (b *Blunders) DumpToCommandLine() {
	fmt.Println("")
	fmt.Println("---------------------")
	fmt.Println("Reported Blunders:")
	for _, blunder := range b.Reported {
		fmt.Println(blunder.Error())
	}
	fmt.Println("")
	fmt.Println("---------------------")
	fmt.Println("Self Blunders:")
	for _, self_blunder := range b.selfBlunders {
		fmt.Println(self_blunder.Error())
	}
	fmt.Println("")

}