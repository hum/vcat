package format

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hum/vcat/pkg/types"
)

// TranscriptToJSON handles the creation of a JSON structure from the given transcript.
// If passed in `prettyFormat` as true, the JSON structure will be whitespace indented in the byte slice.
func TranscriptToJSON(t *types.Transcript, prettyFormat bool) ([]byte, error) {
	if prettyFormat {
		return json.MarshalIndent(t, "", "  ")
	}
	return json.Marshal(t)
}

// TranscriptToCSV takes in a transcript and turns it into a byte slice representation of the CSV.
func TranscriptToCSV(t *types.Transcript) ([]byte, error) {
	var (
		csvData   strings.Builder
		csvWriter = csv.NewWriter(&csvData)
		csvHeader = []string{"start", "end", "duration", "text"}
	)
	if err := csvWriter.Write(csvHeader); err != nil {
		return nil, err
	}

	for _, part := range t.Text {
		var row = []string{
			part.Start,
			part.End,
			fmt.Sprintf("%.2f", part.Duration),
			part.Text,
		}
		csvWriter.Write(row)
	}
	csvWriter.Flush()

	// Make sure Flush did not return an error
	if err := csvWriter.Error(); err != nil {
		return nil, err
	}
	return []byte(csvData.String()), nil
}
