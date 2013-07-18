package raw_data

import (
	"encoding/json"
	"bytes"
	"log"
)

type StripTokenJson struct {
	Direction	string
	Separator	string	
	Number		string
}

func NewStripTokenJson(msg string) (strip_token_json *StripTokenJson) {
	if len(msg) == 0 {
		log.Println("Error contructing StripTokenJson from empty json string.")
		return nil
	}
	
	err := json.Unmarshal([]byte(msg), &strip_token_json)
	
	if err != nil {
		log.Println("Error unmarshalling StripTokenJson from json string.")
		return nil
	}
	
	return strip_token_json
}

func (this *StripTokenJson) Cmp(other *StripTokenJson) (equ bool) {
	if this.Direction != other.Direction {
		return false
	} else if this.Separator != other.Direction {
		return false
	} else if this.Number != other.Number {
		return false
	}
	
	return true
}

func (this *StripTokenJson) String() (out string) {
	var buffer bytes.Buffer
	buffer.WriteString("Direction: " + this.Direction)
	buffer.WriteString(" Separator: " + this.Separator)
	buffer.WriteString(" Number: " + this.Number)
	
	return buffer.String()
}