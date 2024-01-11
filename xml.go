package vcat

import "encoding/xml"

// ParseTranscriptFromXml takes the raw body of the response and turns it into a
// valid struct representation of the transcript.
func ParseTranscriptFromXml(b []byte) (Transcript, error) {
	var transcript Transcript
	err := xml.Unmarshal(b, &transcript)
	return transcript, err
}
