package raw_data

import (
	"log"
	"encoding/json"
	"bytes"
)

type CurroMsgJson struct {
	Protocol		string
	Version		string
	MessageType	string
	Data			map[string]interface{}
}

func NewCurroMsgJson(msg string) (curro_msg_json *CurroMsgJson) {
	if len(msg) == 0 {
		log.Println("Error contructing CurroMsgJson from empty json string.")
		return nil
	}
	
	err := json.Unmarshal([]byte(msg), &curro_msg_json)
	
	if err != nil {
		log.Println("Error unmarshalling CurroMsgJson from json string.")
		return nil
	}
	
	return curro_msg_json
}

func (this *CurroMsgJson) Cmp(other CurroMsgJson) (equ bool) {
	if this.Protocol != other.Protocol {
		return false
	} else if this.Version != other.Version {
		return false
	} else if this.MessageType != other.MessageType {
		return false
	}
	
	return true
}

func (this *CurroMsgJson) String() (out string) {
	var buffer bytes.Buffer
	buffer.WriteString("Protocol: " + this.Protocol)
	buffer.WriteString(" Version: " + this.Version)
	buffer.WriteString(" MessageType: " + this.MessageType)
	
	return buffer.String()
}

