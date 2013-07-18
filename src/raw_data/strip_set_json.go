package raw_data

import (
	"log"
	"encoding/json"
	"bytes"
)

type StripSetJson struct {
	From		string
	To		string
}

func NewStipSetJson(msg string) (strip_set_json *StripSetJson) {
	if len(msg) == 0 {
		log.Println("Error contructing StripSetJson from empty json string.")
		return nil
	}
	
	err := json.Unmarshal([]byte(msg), &strip_set_json)
	
	if err != nil {
		log.Println("Error unmarshalling StripSetJson from json string.")
		return nil
	}
	
	return strip_set_json
}

func (this *StripSetJson) Cmp(other *StripSetJson) (equ bool) {
	if this.From != other.From {
		return false
	} else if this.To != other.To {
		return false
	}
	
	return true
}

func (this *StripSetJson) String() (out string) {
	var buffer bytes.Buffer
	buffer.WriteString("From: " + this.From)
	buffer.WriteString(" To: " + this.To)
	
	return buffer.String()
}
