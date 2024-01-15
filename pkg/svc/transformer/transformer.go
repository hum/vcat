package transformer

import (
	"encoding/json"
	"encoding/xml"
	"strings"

	"github.com/hum/vcat/pkg/types"
	"github.com/mitchellh/mapstructure"
)

// GetCaptionsFromInitialResponse transforms the provided HTML body into a struct,
// which represents the available captions for the given YouTube video.
//
// It receives the initial YouTube response as a byte slice.
//
// @TODO: This is a very naive function which does not handle any edge-cases.
func GetCaptionsFromInitialHttpResponse(b []byte) (types.Captions, error) {
	dataStr := string(b)

	// Unfortunately the response is in HTML, so we parse it as a string
	// and only load the necessary parts as a valid JSON.
	parts := strings.Split(dataStr, "\"captions\":")
	parts = strings.Split(parts[1], ",\"videoDetails\"")

	var jsonString map[string]interface{}
	json.Unmarshal([]byte(parts[0]), &jsonString)

	var captions types.Captions
	err := mapstructure.Decode(jsonString, &captions)
	return captions, err
}

// GetTranscriptFromXMLResponse transforms the provided XML body into a struct,
// which represents the available transcript for the given YouTube video.
//
// It receives the timedtext (https://www.youtube.com/api/timedtext) response as a byte slice.
func GetTranscriptFromXMLResponse(b []byte) (*types.Transcript, error) {
	var transcript types.Transcript
	err := xml.Unmarshal(b, &transcript)
	return &transcript, err
}
