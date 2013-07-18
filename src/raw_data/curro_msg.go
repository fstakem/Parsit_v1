package raw_data

import (
	"log"
	"bytes"
)

type CurroMsg struct {
	Protocol		string
	Version		string
	MessageType	string
	Data			*map[string]interface{}
}

func NewCurroMsg(curro_msg_json *CurroMsgJson) (curro_msg *CurroMsg) {
	if curro_msg_json == nil {
		log.Println("Error contructing CurroMsg from nil object.")
		return nil
	}
	
	// Protocol
	curro_msg.Protocol = curro_msg_json.Protocol
	
	// Version
	curro_msg.Version = curro_msg_json.Version
	
	// MessageType
	curro_msg.MessageType = curro_msg_json.MessageType
	
	// Data
	curro_msg.Data = &curro_msg_json.Data
	
	return curro_msg
}

func (this *CurroMsg) Cmp(other *CurroMsg) (equ bool) {
	if this.Protocol != other.Protocol {
		return false
	} else if this.Version != other.Version {
		return false
	} else if this.MessageType != other.MessageType {
		return false
	}
	
	return true
}

func (this *CurroMsg) String() (out string) {
	var buffer bytes.Buffer
	buffer.WriteString("Protocol: " + this.Protocol)
	buffer.WriteString(" Version: " + this.Version)
	buffer.WriteString(" MessageType: " + this.MessageType)
	
	return buffer.String()
}

