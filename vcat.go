package main

import (
	"flag"
	"fmt"
	"log/slog"
	"os"

	"github.com/hum/vcat/pkg/format"
	"github.com/hum/vcat/pkg/svc"
	"github.com/hum/vcat/pkg/types"
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
	flag.StringVar(&videoURL, "url", "", "url to the video to get transcription from")
	flag.StringVar(&videoURL, "u", "", "url to the video to get transcription from")
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

	var svc = svc.NewTranscriptSvc()

	if showLanguageCodes {
		languages, err := ListLanguages(svc, videoURL)
		if err != nil {
			slog.Error("cannot list languages, err", err)
			os.Exit(1)
		}
		fmt.Println(languages)
		os.Exit(0)
	}

	transcript, err := GetTranscription(svc, videoURL, language)
	if err != nil {
		slog.Error("cannot get transcription", "err", err, "url", videoURL, "language", language)
		os.Exit(1)
	}

	rr, err := FormatTranscriptToByteSlice(transcript, filetypeFormat)
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

func GetTranscription(svc *svc.TranscriptSvc, url string, language string) (*types.Transcript, error) {
	return svc.GetTranscript(url, language)
}

func ListLanguages(svc *svc.TranscriptSvc, url string) ([]types.AvailableLanguage, error) {
	return svc.GetLanguageCodes(videoURL)
}

func FormatTranscriptToByteSlice(t *types.Transcript, ftype string) ([]byte, error) {
	var (
		result []byte
		err    error
	)
	switch filetypeFormat {
	case "json":
		result, err = format.TranscriptToJSON(t, prettyFormat)
		if err != nil {
			return nil, err
		}
	case "csv":
		result, err = format.TranscriptToCSV(t)
		if err != nil {
			return nil, err
		}
	default:
		return nil, fmt.Errorf("unsupported file type=%s", ftype)
	}
	return result, nil
}
