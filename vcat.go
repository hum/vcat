package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/hum/vcat/pkg/svc"
	"github.com/hum/vcat/pkg/types"
)

var (
	videoURL          string
	language          string
	prettyPrint       bool
	showLanguageCodes bool
)

func main() {
	flag.StringVar(&videoURL, "url", "", "url to the video to get transcription from")
	flag.StringVar(&videoURL, "u", "", "url to the video to get transcription from")
	flag.StringVar(&language, "language", "en", "fetch captions in different languages")
	flag.BoolVar(&showLanguageCodes, "l", false, "show a list of available language codes")
	flag.BoolVar(&showLanguageCodes, "list", false, "show a list of available language codes")
	flag.BoolVar(&prettyPrint, "pretty", false, "pretty print the JSON to the CLI")
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

	if prettyPrint {
		fmt.Println(StringIdentStruct(transcript))
	} else {
		fmt.Println(StringStruct(transcript))
	}

	os.Exit(0)
}

func GetTranscription(svc *svc.TranscriptSvc, url string, language string) (*types.Transcript, error) {
	return svc.GetTranscript(url, language)
}

func ListLanguages(svc *svc.TranscriptSvc, url string) ([]types.AvailableLanguage, error) {
	return svc.GetLanguageCodes(videoURL)
}
