package vcat

import (
	"encoding/json"
	"encoding/xml"
	"strconv"
	"strings"
	"time"

	"github.com/mitchellh/mapstructure"
)

// getRawVideoDetailFromInitialResponse transforms the provided HTML body into a struct,
// which represents the available captions, and other metadata, of the video.
//
// It receives the initial HTML YouTube response as a byte slice.
//
// @TODO: This is a very naive function
func getRawVideoDetailFromInitialHttpResponse(b []byte) (*rawVideoDetail, error) {
	dataStr := string(b)

	// Unfortunately the response is in HTML, so we parse it as a string
	// and only load the necessary parts as a valid JSON.
	parts := strings.Split(dataStr, "\"captions\":")
	if len(parts) < 2 {
		return nil, ErrCaptionsNotFound
	}
	parts = strings.Split(parts[1], ",\"videoDetails\"")

	// We also care about the title, thumbnail, etc.
	metadataParts := strings.Split(parts[1], ",\"playerConfig\"")

	var (
		rawCaptions = parts[0]
		rawMetadata = metadataParts[0][1:] // Remove the ":" prefix of the string, so that it is a valid JSON

		captionsMap map[string]interface{}
		metadataMap map[string]interface{}
	)

	if err := json.Unmarshal([]byte(rawCaptions), &captionsMap); err != nil {
		return nil, ErrCaptionsNotFound
	}
	var captions captions
	if err := mapstructure.Decode(captionsMap, &captions); err != nil {
		return nil, ErrCaptionsNotFound
	}
	if err := json.Unmarshal([]byte(rawMetadata), &metadataMap); err != nil {
		return nil, ErrCaptionsNotFound
	}
	var metadata VideoMetadata
	if err := mapstructure.Decode(metadataMap, &metadata); err != nil {
		return nil, ErrCaptionsNotFound
	}
	return &rawVideoDetail{metadata: &metadata, captions: captions}, nil
}

// GetTranscriptFromXMLResponse transforms the provided XML body into a struct,
// which represents the available transcript for the given YouTube video.
//
// It receives the timedtext (https://www.youtube.com/api/timedtext) response as a byte slice.
func getTranscriptFromXMLResponse(b []byte) (*transcript, error) {
	var transcript transcript
	err := xml.Unmarshal(b, &transcript)
	if err != nil {
		return nil, ErrTranscriptNotFound
	}
	unescapeCharactersFromText(&transcript)
	return parseStartTimeEndTimeTimestamps(transcript)
}

// Parses the raw transcript into a more human-readable form with normalised datetime representations.
// {"start": "0.0", "duration": "1.0", "text": "hello"} => {"start": "00:00:00", "end": "00:00:01", "duration": 1.0, "text": "hello"}
func parseStartTimeEndTimeTimestamps(t transcript) (*transcript, error) {
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
func unescapeCharactersFromText(t *transcript) {
	for i, txt := range t.Text {
		t.Text[i].Text = strings.ReplaceAll(txt.Text, "\u0026#39;", "'")
		t.Text[i].Text = strings.ReplaceAll(txt.Text, "\u0026amp;#39;", "'")
	}
}
