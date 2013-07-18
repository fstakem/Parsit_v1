package raw_data

import (
	"log"
	"encoding/json"
	"bytes"
)

type LogRulesJson struct {
	LogKey			string
	LastModified		string
	Rules			map[string]interface{}
}

func NewLogRulesJson(msg string) (log_rules_json *LogRulesJson) {
	if len(msg) == 0 {
		log.Println("Error contructing LogRulesJson from empty json string.")
		return nil
	}
	
	err := json.Unmarshal([]byte(msg), &log_rules_json)
	
	if err != nil {
		log.Println("Error unmarshalling LogRulesJson from json string.")
		return nil
	}
	
	return log_rules_json
}

func (this *LogRulesJson) Cmp(other *LogRulesJson) (equ bool) {
	if this.LogKey != other.LogKey {
		return false
	} else if this.LastModified != other.LastModified {
		return false
	}
	
	return true	
}

func (this *LogRulesJson) String() (out string) {
	var buffer bytes.Buffer
	buffer.WriteString("LogKey: " + this.LogKey)
	buffer.WriteString(" LastModified: " + this.LastModified)
	
	return buffer.String()
}

