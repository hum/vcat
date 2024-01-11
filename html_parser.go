package vcat

import (
	"encoding/json"
	"strings"

	"github.com/mitchellh/mapstructure"
)

// GetCaptionsFromRawHtml is an ad-hoc function to parse the provided HTML body
// into a struct, which represents the available captions for a given YouTube video.
//
// @TODO: This is a very naive function which does not handle any edge-cases.
func GetCaptionsFromRawHtml(b []byte) (Captions, error) {
	dataStr := string(b)

	// Magic which will most likely break in the future.
	// Unfortunately the response is in HTML, so we parse it as a string
	// and only load the necessary parts as a valid JSON.
	parts := strings.Split(dataStr, "\"captions\":")
	parts = strings.Split(parts[1], ",\"videoDetails\"")

	var jsonString map[string]interface{}
	json.Unmarshal([]byte(parts[0]), &jsonString)

	var captions Captions
	err := mapstructure.Decode(jsonString, &captions)
	return captions, err
}
