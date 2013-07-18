package raw_data

import (
	"log"
	"encoding/json"
	"bytes"
)

type StripLengthJson struct {
	From		string
	To		string
}

func NewStripLengthJson(msg string) (strip_length_json *StripLengthJson) {
	if len(msg) == 0 {
		log.Println("Error contructing StripLengthJson from empty json string.")
		return nil
	}
	
	err := json.Unmarshal([]byte(msg), &strip_length_json)
	
	if err != nil {
		log.Println("Error unmarshalling StripLengthJson from json string.")
		return nil
	}
	
	return strip_length_json
}

func (this *StripLengthJson) Cmp(other *StripLengthJson) (equ bool) {
	if this.From != other.From {
		return false
	} else if this.To != other.To {
		return false
	}
	
	return true
}

func (this *StripLengthJson) String() (out string) {
	var buffer bytes.Buffer
	buffer.WriteString("From: " + this.From)
	buffer.WriteString(" To: " + this.To)
	
	return buffer.String()
}
