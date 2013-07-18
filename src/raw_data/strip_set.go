package raw_data

import (
	"log"
	"bytes"
)

type StripSet struct {
	From		string
	To		string
}

func NewStripSet(strip_set_json *StripSetJson) (strip_set *StripSet) {
	if strip_set_json == nil {
		log.Println("Error contructing StripSet from nil object.")
		return nil
	}
	
	// From
	strip_set.From = strip_set_json.From
	
	// To
	strip_set.To = strip_set_json.To
	
	return strip_set
}

func (this *StripSet) Cmp(other *StripSet) (equ bool) {
	if this.From != other.From {
		return false
	} else if this.To != other.To {
		return false
	}
	
	return true
}

func (this *StripSet) String() (out string) {
	var buffer bytes.Buffer
	buffer.WriteString("From: " + this.From)
	buffer.WriteString(" To: " + this.To)
	
	return buffer.String()
}

