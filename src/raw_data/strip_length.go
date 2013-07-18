package raw_data

import (
	"log"
	"strconv"
	"bytes"
)

type StripLength struct {
	From		int
	To		int
}

func NewStripLength(strip_length_json *StripLengthJson) (strip_length *StripLength) {
	if strip_length_json == nil {
		log.Println("Error contructing StripLength from nil object.")
		return nil
	}
	
	// From
	from, err := strconv.Atoi(strip_length_json.From)
	if err != nil {
		log.Println("StripLength.From not properly parsed.")
		return nil
	}
	strip_length.From = from
	
	// To
	to, err := strconv.Atoi(strip_length_json.To)
	if err != nil {
		log.Println("StripLength.To not properly parsed.")
		return nil
	}
	strip_length.To = to
	
	return strip_length
}

func (this *StripLength) Cmp(other *StripLength) (equ bool) {
	if this.From != other.From {
		return false
	} else if this.To != other.To {
		return false
	}
	
	return true
}

func (this *StripLength) String() (out string) {
	var buffer bytes.Buffer
	buffer.WriteString("From: " + strconv.Itoa(this.From))
	buffer.WriteString(" To: " + strconv.Itoa(this.To))
	
	return buffer.String()
}

