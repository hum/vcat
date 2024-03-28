package main

import (
	"encoding/csv"
	"encoding/json"
	"flag"
	"fmt"
	"log/slog"
	"os"
	"strings"

	"github.com/hum/vcat"
)

var (
	videoURL          string
	language          string
	outputPath        string
	filetypeFormat    string
	prettyFormat      bool
	showLanguageCodes bool
)

func main() {
	flag.StringVar(&videoURL, "url", "", "the url to the video to get a transcription from")
	flag.StringVar(&videoURL, "u", "", "the url to the video to get a transcription from")
	flag.StringVar(&language, "language", "en", "fetch captions in different languages")
	flag.StringVar(&filetypeFormat, "format", "json", "specify which format to use for the data")
	flag.StringVar(&outputPath, "o", "", "output path of the file")
	flag.BoolVar(&showLanguageCodes, "l", false, "show a list of available language codes")
	flag.BoolVar(&showLanguageCodes, "list", false, "show a list of available language codes")
	flag.BoolVar(&prettyFormat, "pretty", false, "pretty print the JSON to the CLI")
	flag.Parse()

	if videoURL == "" {
		flag.Usage()
		os.Exit(1)
	}

	if showLanguageCodes {
		languages, err := vcat.GetAvailableCaptionLanguages(videoURL)
		if err != nil {
			slog.Error("cannot list languages, err", err)
			os.Exit(1)
		}
		fmt.Println(languages)
		os.Exit(0)
	}

	transcript, err := vcat.GetVideoTranscript(videoURL, language)
	if err != nil {
		slog.Error("cannot get transcription", "err", err, "url", videoURL, "language", language)
		os.Exit(1)
	}

	rr, err := formatTranscriptToByteSlice(transcript, filetypeFormat, prettyFormat)
	if err != nil {
		slog.Error("cannot format transcript")
		os.Exit(1)
	}

	if outputPath != "" {
		err := os.WriteFile(outputPath, rr, 0644)
		if err != nil {
			slog.Error("cannot write transcription to file", "err", err)
			os.Exit(1)
		}
	} else {
		fmt.Println(string(rr))
	}
	os.Exit(0)
}

func formatTranscriptToByteSlice(t []vcat.TextChunk, ftype string, prettyFormat bool) ([]byte, error) {
	var (
		result []byte
		err    error
	)
	switch ftype {
	case "json":
		result, err = transcriptToJSON(t, prettyFormat)
		if err != nil {
			return nil, err
		}
	case "csv":
		result, err = transcriptToCSV(t)
		if err != nil {
			return nil, err
		}
	default:
		return nil, fmt.Errorf("unsupported file type=%s", ftype)
	}
	return result, nil
}

func transcriptToJSON(t []vcat.TextChunk, prettyFormat bool) ([]byte, error) {
	if prettyFormat {
		return json.MarshalIndent(t, "", "  ")
	}
	return json.Marshal(t)
}

func transcriptToCSV(t []vcat.TextChunk) ([]byte, error) {
	var (
		csvData   strings.Builder
		csvWriter = csv.NewWriter(&csvData)
		csvHeader = []string{"start", "end", "duration", "text"}
	)
	if err := csvWriter.Write(csvHeader); err != nil {
		return nil, err
	}

	for _, part := range t {
		var row = []string{
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
