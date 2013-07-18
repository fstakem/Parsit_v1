package raw_data

import (
	"log"
	"encoding/json"
	"bytes"
)

type StripPatternJson struct {
	Direction	string
	Pattern		string
}

func NewStripPatternJson(msg string) (strip_pattern_json *StripPatternJson) {
	if len(msg) == 0 {
		log.Println("Error contructing StripPatternJson from empty json string.")
		return nil
	}
	
	err := json.Unmarshal([]byte(msg), &strip_pattern_json)
	
	if err != nil {
		log.Println("Error unmarshalling StripPatternJson from json string.")
		return nil
	}
	
	return strip_pattern_json
}

func (this *StripPatternJson) Cmp(other *StripPatternJson) (equ bool) {
	if this.Direction != other.Direction {
		return false
	} else if this.Pattern != other.Pattern {
		return false
	}
	
	return true
}

func (this *StripPatternJson) String() (out string) {
	var buffer bytes.Buffer
	buffer.WriteString("Direction: " + this.Direction)
	buffer.WriteString(" Pattern: " + this.Pattern)
	
	return buffer.String()
}

