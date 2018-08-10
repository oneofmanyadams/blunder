package blunders

import (
	"time"
	"io"
	"os"
	"strconv"
	"fmt"
)


//////////////////////////////////////////////////////////////////
// BlunderBus Code
//////////////////////////////////////////////////////////////////

type BlunderBus struct {
	Blunders []Blunder
	HasFatal bool
	ExitOnFatal bool
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

func (bb *BlunderBus) New(code string, message string) Blunder {
	return bb.newBase(code, message, false, time.Now())
}

func (bb *BlunderBus) NewFatal(code string, message string) Blunder {
	return bb.newBase(code, message, true, time.Now())
}

func (bb *BlunderBus) newBase(code string, message string, fatal bool, b_time time.Time) (new_blunder Blunder) {
	new_blunder = NewBlunder(code, message, fatal, b_time)
	bb.Blunders = append(bb.Blunders, new_blunder)
	if fatal {
		bb.HasFatal = true
	}
	if bb.HasFatal && bb.ExitOnFatal {
		fmt.Println("")
		fmt.Println("Fatal blunder encountered.")
		bb.LogDump()
		fmt.Println("")
		os.Exit(1)
	}
	return
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
// BlunderBus Consolidation Methods
//////////////////////////////////////////////////////////////////
func (bb *BlunderBus) IncludeBlundersFrom(imported_bb *BlunderBus) {
	for _, imported_blunder := range imported_bb.Blunders {
		bb.newBase(imported_blunder.Code, imported_blunder.Message, imported_blunder.Fatal, imported_blunder.Time)
	}
}

//////////////////////////////////////////////////////////////////
// BlunderBus Utility Methods
//////////////////////////////////////////////////////////////////

func (bb BlunderBus) BlunderSliceAsString(blunder_slice []Blunder) (blunder_string string) {
	for _, blndr := range blunder_slice {
		blunder_string = blunder_string + blndr.Error() + "\n"
	}
	return
}

func (bb BlunderBus) LogTo(writer io.Writer) {
	all_blunders := []byte(bb.BlunderSliceAsString(bb.Blunders))
	writer.Write(all_blunders)
}

func (bb BlunderBus) LogDump() {
	log_file, log_file_error := os.Create("error_log_"+strconv.Itoa(int(time.Now().Unix()))+".log")
	if log_file_error != nil {
		bb.New("FILECREATE", log_file_error.Error())
		fmt.Println("Wow, you really screwed up. I can't even create the blunder log.")
		fmt.Println("Dumping blunders to Stdout in 5 seconds...")
		fmt.Println("")
		time.Sleep(5 * time.Second)
		bb.LogTo(os.Stdout)
	} else {
		fmt.Println("Dumping errors to log...")
		bb.LogTo(log_file)
		log_file.Close()
	}
}