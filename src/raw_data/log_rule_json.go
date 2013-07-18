package raw_data

import (
	"log"
	"encoding/json"
	"bytes"
)

type LogRuleJson struct {
	RuleType			string
	RuleParameters	map[string]interface{}
}

func NewLogRuleJson(msg string) (log_rule_json *LogRuleJson) {
	if len(msg) == 0 {
		log.Println("Error contructing LogRuleJson from empty json string.")
		return nil
	}
	
	err := json.Unmarshal([]byte(msg), &log_rule_json)
	
	if err != nil {
		log.Println("Error unmarshalling LogRuleJson from json string.")
		return nil
	}
	
	return log_rule_json
}

func (this *LogRuleJson) Cmp(other *LogRuleJson) (equ bool) {
	if this.RuleType != other.RuleType {
		return false
	}
	
	return true
}

func (this *LogRuleJson) String() (out string) {
	var buffer bytes.Buffer
	buffer.WriteString("RuleType: " + this.RuleType)
	
	return buffer.String()
}

