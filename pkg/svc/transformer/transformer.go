package transformer

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"strconv"
	"strings"
	"time"

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
	if len(parts) < 2 {
		return types.Captions{}, fmt.Errorf("could not find captions")
	}
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
	if err != nil {
		return nil, err
	}
	unescapeCharactersFromText(&transcript)
	return parseStartTimeEndTimeTimestamps(transcript)
}

// Parses the raw transcript into a more human-readable form with normalised datetime representations.
// {"start": "0.0", "duration": "1.0", "text": "hello"} => {"start": "00:00:00", "end": "00:00:01", "duration": 1.0, "text": "hello"}
func parseStartTimeEndTimeTimestamps(t types.Transcript) (*types.Transcript, error) {
	for i, item := range t.Text {
		startTimeFloat, err := strconv.ParseFloat(item.Start, 64)
		if err != nil {
			return nil, err
		}
		var (
			startOffset = time.Duration(startTimeFloat) * time.Second
			duration    = time.Duration(item.Duration) * time.Second
			endOffset   = time.Duration(startOffset + duration)

			startTime time.Time
			endTime   time.Time
		)

		startTime = startTime.Add(startOffset)
		endTime = endTime.Add(endOffset)

		t.Text[i].Start = startTime.Format("15:04:05")
		t.Text[i].End = endTime.Format("15:04:05")
	}
	return &t, nil
}

// unescapeCharactersFromText removes unecessary character encoding possibly present in the raw text field
func unescapeCharactersFromText(t *types.Transcript) {
	for i, txt := range t.Text {
		t.Text[i].Text = strings.ReplaceAll(txt.Text, "\u0026#39;", "'")
		t.Text[i].Text = strings.ReplaceAll(txt.Text, "\u0026amp;#39;", "'")
	}
}
