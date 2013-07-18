package raw_data

import (
	"log"
	"bytes"
)

type LogRule struct {
	RuleType			string
	RuleParameters	*map[string]interface{}
}

func NewLogRule(log_rule_json *LogRuleJson) (log_rule *LogRule) {
	if log_rule_json == nil {
		log.Println("Error contructing LogRule from nil object.")
		return nil
	}
	
	// RuleType
	log_rule.RuleType = log_rule_json.RuleType
	
	// RuleParameters
	log_rule.RuleParameters = &log_rule_json.RuleParameters
	
	return log_rule
}

func (this *LogRule) Cmp(other *LogRule) (equ bool) {
	if this.RuleType != other.RuleType {
		return false
	}
	
	return true
}

func (this *LogRule) String() (out string) {
	var buffer bytes.Buffer
	buffer.WriteString("RuleType: " + this.RuleType)
	
	return buffer.String()
}
