package vcat

import (
	"encoding/json"
)

// StringIdentStruct helps format a struct into a pretty JSON. Uses tabs.
func StringIdentStruct(strct any) string {
	s, _ := json.MarshalIndent(strct, "", "\t")
	return string(s)
}
