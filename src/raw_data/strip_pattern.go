package raw_data

import (
	"bytes"
	"log"
)

type StripPattern struct {
	Direction	ParseDirection
	Pattern		string
}

func NewStripPattern(strip_pattern_json *StripPatternJson) (strip_pattern *StripPattern) {
	if strip_pattern_json == nil {
		log.Println("Error contructing StripPattern from nil object.")
		return nil
	}
	
	// Direction
	if strip_pattern_json.Direction == "forward" {
		strip_pattern.Direction = ParseForward
	} else if strip_pattern_json.Direction == "backward" {
		strip_pattern.Direction = ParseBackward
	} else {
		log.Println("Error parsing Direction from StripPatternJson.")
	}
	
	// Pattern
	strip_pattern.Pattern = strip_pattern_json.Pattern
	
	return strip_pattern
}

func (this *StripPattern) Cmp(other *StripPattern) (equ bool) {
	if this.Direction != other.Direction {
		return false
	} else if this.Pattern != other.Pattern {
		return false
	}
	
	return true
}

func (this *StripPattern) String() (out string) {
	var buffer bytes.Buffer
	var direction = ""
	if this.Direction == ParseForward {
		direction = "forward"
	} else if this.Direction == ParseBackward {
		direction = "backward"
	} 
	
	buffer.WriteString("Direction: " + direction)
	buffer.WriteString(" Pattern: " + this.Pattern)
	
	return buffer.String()
}

