package raw_data

import (
	"bytes"
	"log"
	"strconv"
)

type StripToken struct {
	Direction	ParseDirection
	Separator	string
	Number		uint64
}

func NewStripToken(strip_token_json *StripTokenJson) (strip_token *StripToken) {
	if strip_token_json == nil {
		log.Println("Error contructing StripToken from nil object.")
		return nil
	}
	
	// Direction
	if strip_token_json.Direction == "forward" {
		strip_token.Direction = ParseForward
	} else if strip_token_json.Direction == "backward" {
		strip_token.Direction = ParseBackward
	} else {
		log.Println("Error parsing Direction from StripTokenJson.")
	}
	
	// Serarator
	strip_token.Separator = strip_token_json.Separator
	
	// Number
	number, err := strconv.ParseUint(strip_token_json.Number, 10, 64)
	if err != nil {
		log.Println("Error parsing Number from StripTokenJson.")
		return nil
	}
	strip_token.Number = number
	
	return strip_token
}

func (this *StripToken) Cmp(other *StripToken) (equ bool) {
	if this.Direction != other.Direction {
		return false
	} else if this.Separator != other.Separator {
		return false
	} else if this.Number != other.Number {
		return false
	}
	
	return true
}

func (this *StripToken) String() (out string) {
	var buffer bytes.Buffer
	var direction = ""
	if this.Direction == ParseForward {
		direction = "forward"
	} else if this.Direction == ParseBackward {
		direction = "backward"
	} 
	
	buffer.WriteString("Direction: " + direction)
	buffer.WriteString(" Separator: " + this.Separator)
	buffer.WriteString(" Number: " + strconv.FormatUint(this.Number, 10))
	
	return buffer.String()
}

