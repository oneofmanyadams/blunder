package blunders

import "testing"

//////////////////////////////////////////////////////////////////
// Examples
//////////////////////////////////////////////////////////////////

//////////////////////////////////////////////////////////////////
// Initialization Tests
//////////////////////////////////////////////////////////////////

func TestNewBlunders(t *testing.T) {
	type blunders_tester struct {
		id string
		test_type string
	}
	tables := []blunders_tester{
			{"test", "One word identifier."},
			{"t", "Single letter identifier."},
			{"", "No identifier."},
			{"Test Instance", "Identifier with space."},
			{"TestInstance ", "Identifier with trailing space."},
			{" TestInstance", "Identifier with leading space."}}

	for _, bt := range tables {
		new_blunders := NewBlunders(bt.id)

		if new_blunders.Identifier != bt.id {
			t.Errorf("Blunder identifier does not match provided id in "+bt.test_type+" test.")
		}
		if len(new_blunders.Codes) > 1 {
			t.Errorf("Too many blunder codes created on initialization in "+bt.test_type+" test.")
		}
		if len(new_blunders.Codes) < 1 {
			t.Errorf("Not enough blunder codes created on initialization in "+bt.test_type+" test.")
		}
		if new_blunders.Codes[0] != "SelfBlunder" {
			t.Errorf("SelfBlunder code type not set as member 0 in Codes map on initialization in "+bt.test_type+" test.")
		}
		if len(new_blunders.Reported) > 0 {
			t.Errorf("Instance created with reported blunders on initialization in "+bt.test_type+" test.")
		}
		if len(new_blunders.selfBlunders) > 0 {
			t.Errorf("Instance created with reported self blunders on initialization in "+bt.test_type+" test.")
		}
		if new_blunders.Codes[0] != "SelfBlunder" {
			t.Errorf("SelfBlunder code type not set as member 0 in Codes map on initialization in "+bt.test_type+" test.")
		}

	}
}

//////////////////////////////////////////////////////////////////
// Code Tests
//////////////////////////////////////////////////////////////////

func TestAddCode(t *testing.T) {
	type code_test struct {
		number int
		name string
		should_pass bool
		fail_message string
	}
	tables := []code_test{
			{0, "Zero", false, "Add Code accpeted a Code number reserved for SelfBlunders (Code ID: 0)."},
			{1, "", true, "Add Code did not add code with no name."},
			{1, "One", true, ""},
			{20, "Twenty", true, "Did not accept non-sequential Code number."}}

	for _, bt := range tables {
		blndrs := NewBlunders("test")
		added := blndrs.AddCode(bt.number, bt.name)
		
		if bt.should_pass {
			if added == false {
				t.Errorf("Blunder code add failed (Method returned false). "+bt.fail_message)
			}
			if saved_code_name, code_num_exists := blndrs.Codes[bt.number]; code_num_exists {
				if (saved_code_name != bt.name) {
					t.Errorf("Code name recorded not the same as name provided. "+bt.fail_message)
				}
			} else {
				t.Errorf("Method reported successful code addition, but code number does not exist. "+bt.fail_message)
			}

		} else {
			if added == true {
				t.Errorf("Blunder code add succeeded when should have failed (Method returned true). "+bt.fail_message)
			}
		}

		if blndrs.Codes[0] != "SelfBlunder" {
			t.Errorf("Self Blunder code was overwritten. "+bt.fail_message)
		}
	}
}

func TestAddCode_duplicate_number(t *testing.T) {
	blndrs := NewBlunders("test")
	added := blndrs.AddCode(1, "one")
	added2 := blndrs.AddCode(1, "two")

	if added == false {
		t.Errorf("Unable to create basic Blunder.")
	}
	if added2 != false {
		t.Errorf("Successfully added duplicate blunder number (should have failed).")
	}
	if len(blndrs.selfBlunders) < 1 {
		t.Errorf("Did not record a SelfBlunder when a blunder number was duplicated.")
	}
	if blndrs.Codes[1] != "one" {
		t.Errorf("Original blunder name overwritten by duplication attempt.")
	}

}

func TestAddCode_duplicate_name(t *testing.T) {
	blndrs := NewBlunders("test")
	added := blndrs.AddCode(1, "one")
	added2 := blndrs.AddCode(2, "one")

	if added == false {
		t.Errorf("Unable to create basic Blunder.")
	}
	if added2 == true {
		t.Errorf("Successfully added duplicate blunder name (should have failed).")
	}
	if len(blndrs.selfBlunders) < 1 {
		t.Errorf("Did not record a SelfBlunder when a blunder name was duplicated.")
	}
	if _, duplicate_saved := blndrs.Codes[2]; duplicate_saved {
		t.Errorf("duplicate name record saved.")
	}
}

func TestRemoveCode(t *testing.T) {
	blndrs := NewBlunders("test")
	blndrs.AddCode(1, "one")

	if blndrs.RemoveCode(1) {
		t.Errorf("Method RemoveCode should not ever work!!")
	}
}

//////////////////////////////////////////////////////////////////
// Blunder Creation Tests
//////////////////////////////////////////////////////////////////

func TestNew(t *testing.T) {
	type code_test struct {
		code int
		code_name string
		message string
		should_pass bool
		fail_message string
	}
	tables := []code_test{
			{1, "One", "Test Blunder", true, "Basic New Blunder failure."},
			{1, "One", "", true, "No Message Blunder not added"},
			{15, "One", "Test Blunder", false, "Able to add a non-existant blunder code"}}

	blndrs := NewBlunders("test")
	blndrs.AddCode(1, "One")

	for _, ct := range tables {
		generated_blunder := blndrs.New(ct.code, ct.message)
		if ct.should_pass {
			if generated_blunder.Code != 1 {
				t.Errorf("Blunder code provided not the same as the one recorded. "+ct.fail_message)
			}
			if generated_blunder.CodeName != "One" {
				t.Errorf("Blunder code_name provided not the same as the one recorded."+ct.fail_message)
			}
			if generated_blunder.Fatal {
				t.Errorf("Blunder recorded FATAL blunder when should have recorded NON-FATAL."+ct.fail_message)
			}
			if generated_blunder.Message != ct.message {
				t.Errorf("Message recorded not the same as one provided. "+ct.fail_message)
			}
		} else {
			if generated_blunder.Code != 0 {
				t.Errorf("Unregisted Code ussage should result in SelfBlunder code being used, but that didn't happen. "+ct.fail_message)
			}
			if generated_blunder.CodeName != "SelfBlunder" {
				t.Errorf("Unregisted Code ussage should result in SelfBlunder code_name being used, but that didn't happen. "+ct.fail_message)
			}
			if generated_blunder.Fatal {
				t.Errorf("Blunder recorded FATAL blunder when should have recorded NON-FATAL."+ct.fail_message)
			}

		}
	}
}

func TestNewFatal(t *testing.T) {
	type code_test struct {
		code int
		code_name string
		message string
		should_pass bool
		fail_message string
	}
	tables := []code_test{
			{1, "One", "Test Blunder", true, "Basic New Blunder failure."},
			{1, "One", "", true, "No Message Blunder not added"},
			{15, "One", "Test Blunder", false, "Able to add a non-existant blunder code"}}

	blndrs := NewBlunders("test")
	blndrs.AddCode(1, "One")

	for _, ct := range tables {
		generated_blunder := blndrs.NewFatal(ct.code, ct.message)
		if ct.should_pass {
			if generated_blunder.Code != 1 {
				t.Errorf("Blunder code provided not the same as the one recorded. "+ct.fail_message)
			}
			if generated_blunder.CodeName != "One" {
				t.Errorf("Blunder code_name provided not the same as the one recorded."+ct.fail_message)
			}
			if !generated_blunder.Fatal {
				t.Errorf("Blunder recorded NON-FATAL blunder when should have recorded FATAL."+ct.fail_message)
			}
			if generated_blunder.Message != ct.message {
				t.Errorf("Message recorded not the same as one provided. "+ct.fail_message)
			}
		} else {
			if generated_blunder.Code != 0 {
				t.Errorf("Unregisted Code ussage should result in SelfBlunder code being used, but that didn't happen. "+ct.fail_message)
			}
			if generated_blunder.CodeName != "SelfBlunder" {
				t.Errorf("Unregisted Code ussage should result in SelfBlunder code_name being used, but that didn't happen. "+ct.fail_message)
			}
			if !generated_blunder.Fatal {
				t.Errorf("Blunder recorded NON-FATAL blunder when should have recorded FATAL."+ct.fail_message)
			}

		}
	}
}

//////////////////////////////////////////////////////////////////
// Information Tests
//////////////////////////////////////////////////////////////////

func TestHasFatal_with_fatals(t *testing.T) {
	blndrs := NewBlunders("test")
	blndrs.AddCode(1, "One")
	blndrs.NewFatal(1, "BOOM!")

	if !blndrs.HasFatal() {
		t.Errorf("Reporting false when there IS a FATAL blunder.")
	}

}

func TestHasFatal_without_fatals(t *testing.T) {
	blndrs := NewBlunders("test")
	blndrs.AddCode(1, "One")
	blndrs.New(1, "poof")

	if blndrs.HasFatal() {
		t.Errorf("Reporting true when there is NOT a FATAL blunder.")
	}

}

func TestNoneFatal_with_fatals(t *testing.T) {
	blndrs := NewBlunders("test")
	blndrs.AddCode(1, "One")
	blndrs.NewFatal(1, "BOOM!")

	if blndrs.NoneFatal() {
		t.Errorf("Reporting true when there IS a FATAL blunder.")
	}
}

func TestNoneFatal_without_fatals(t *testing.T) {
	blndrs := NewBlunders("test")
	blndrs.AddCode(1, "One")
	blndrs.New(1, "poof")

	if !blndrs.NoneFatal() {
		t.Errorf("Reporting false when there is NOT a FATAL blunder.")
	}
}