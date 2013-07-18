package raw_data

import (
	"log"
	"bytes"
)

type LogRules struct {
	LogKey			string
	LastModified		string
	Rules			*map[string]interface{}
}

func NewLogRules(log_rules_json *LogRulesJson) (log_rules *LogRules) {
	if log_rules_json == nil {
		log.Println("Error contructing LogRules from nil object.")
		return nil
	}
	
	// LogKey
	log_rules.LogKey = log_rules_json.LogKey
	
	// LastModified
	log_rules.LastModified = log_rules_json.LastModified
	
	// Rules
	log_rules.Rules = &log_rules_json.Rules
	
	return log_rules
}

func (this *LogRules) Cmp(other *LogRules) (equ bool) {
	if this.LogKey != other.LogKey {
		return false
	} else if this.LastModified != other.LastModified {
		return false
	}
	
	return false
}

func (this *LogRules) String() (out string) {
	var buffer bytes.Buffer
	buffer.WriteString("LogKey: " + this.LogKey)
	buffer.WriteString(" LastModified: " + this.LastModified)
	
	return buffer.String()
}

