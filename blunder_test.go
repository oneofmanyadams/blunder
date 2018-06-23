package blunders

import (
	"testing"
	"strconv"
	"fmt"
)
//////////////////////////////////////////////////////////////////
// Examples
//////////////////////////////////////////////////////////////////

func ExampleNewBlunder() {
	blndr := NewBlunder(1, "Example", false, "This is an example blunder.")
	fmt.Println(blndr.Error())
	// Output: NON-FATAL BLUNDER enountered, CODE: 1 (Example), "This is an example blunder."
}

func ExampleBlunder_Error() {
	blndr := NewBlunder(1, "Example", false, "This is an example blunder.")
	fmt.Println(blndr.Error())
	// Output: NON-FATAL BLUNDER enountered, CODE: 1 (Example), "This is an example blunder."
}

//////////////////////////////////////////////////////////////////
// Tests
//////////////////////////////////////////////////////////////////

func TestNewBlunder(t *testing.T) {
	type blunder_tester struct {
		code int
		code_name string
		fatal_bool bool
		message string
		should_pass bool
		if_fail string
	}
	tables := []blunder_tester{
			{1, "one", false, "test blunder", true, "Simple blunder creation failed."},
			{-1, "one", false, "test blunder", true, "Negative code blunder creation failed."},
			{-1, "one", true, "test blunder", true, "Fatal blunder creation failed."},
	}

	for _, bt := range tables {
		cb := 	NewBlunder(bt.code, bt.code_name, bt.fatal_bool, bt.message)

		if !((
			cb.Code == bt.code &&
			cb.CodeName == bt.code_name &&
			cb.Fatal == bt.fatal_bool &&
			cb.Message == bt.message) && 
			bt.should_pass ) ||
			((
			cb.Code != bt.code ||
			cb.CodeName != bt.code_name ||
			cb.Fatal != bt.fatal_bool ||
			cb.Message != bt.message) &&
			!bt.should_pass ) {
				t.Errorf(bt.if_fail)			
		}
	}
}

func TestError(t *testing.T) {
	type blunder_tester struct {
		code int
		code_name string
		fatal_bool bool
		message string
		should_pass bool
		if_fail string
	}
	tables := []blunder_tester{
		{1, "one", false, "test blunder", true, "Simple blunder Error message failed."},
		{-1, "one", false, "test blunder", true, "Negative code Error message failed."},
		{-1, "one", true, "test blunder", true, "Fatal blunder Error message failed."},
	}

	for _, bt := range tables {
		cb := 	NewBlunder(bt.code, bt.code_name, bt.fatal_bool, bt.message)

		var fatal_string string
		if bt.fatal_bool {
			fatal_string = "FATAL BLUNDER"
		} else {
			fatal_string = "NON-FATAL BLUNDER"
		}

		var code_string string
		code_string = strconv.Itoa(bt.code)

		expected_error_string := fatal_string+" enountered, CODE: "+code_string+" ("+bt.code_name+"), \""+bt.message+"\""

		if !((expected_error_string == cb.Error()) && bt.should_pass ) &&
		!((expected_error_string != cb.Error()) && !bt.should_pass ) {
			t.Errorf(bt.if_fail)			
		}
	}

}