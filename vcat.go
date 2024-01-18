package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/hum/vcat/pkg/format"
	"github.com/hum/vcat/pkg/svc"
	"github.com/hum/vcat/pkg/types"
)

type DataFormat string

const (
	JSON DataFormat = "json"
	CSV  DataFormat = "csv"
)

var (
	videoURL          string
	language          string
	prettyFormat      bool
	showLanguageCodes bool
	outputPath        string
	filetypeFormat    string
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
			panic(err)
		}
		fmt.Println(languages)
		os.Exit(0)
	}

	transcript, err := GetTranscription(svc, videoURL, language)
	if err != nil {
		panic(err)
	}

	var rr []byte
	switch filetypeFormat {
	case "json":
		rr, err = format.TranscriptToJSON(transcript, prettyFormat)
		if err != nil {
			panic(err)
		}
	case "csv":
		rr, err = format.TranscriptToCSV(transcript)
		if err != nil {
			panic(err)
		}
	default:
		panic(fmt.Sprintf("unsupported file type=%s", filetypeFormat))
	}

	if outputPath != "" {
		err := os.WriteFile(outputPath, rr, 0644)
		if err != nil {
			panic(err)
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
