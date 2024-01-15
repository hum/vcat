package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/hum/vcat/pkg/svc"
)

var (
	videoURL          string
	language          string
	showLanguageCodes bool
)

func main() {
	flag.StringVar(&videoURL, "url", "", "url to the video to get transcription from")
	flag.StringVar(&videoURL, "u", "", "url to the video to get transcription from")
	flag.StringVar(&language, "language", "en", "fetch captions in different languages")
	flag.BoolVar(&showLanguageCodes, "l", false, "show a list of available language codes")
	flag.BoolVar(&showLanguageCodes, "list", false, "show a list of available language codes")
	flag.Parse()

	if videoURL == "" {
		flag.Usage()
		os.Exit(1)
	}

	var svc = svc.NewTranscriptSvc()

	if showLanguageCodes {
		err := ListLanguages(svc, videoURL)
		if err != nil {
			panic(err)
		}
		os.Exit(0)
	}

	err := GetTranscription(svc, videoURL, language)
	if err != nil {
		panic(err)
	}
	os.Exit(0)
}

func GetTranscription(svc *svc.TranscriptSvc, url string, language string) error {
	transcript, err := svc.GetTranscript(url, language)
	if err != nil {
		return err
	}
	fmt.Println(StringIdentStruct(transcript))
	return nil
}

func ListLanguages(svc *svc.TranscriptSvc, url string) error {
	languages, err := svc.GetLanguageCodes(videoURL)
	if err != nil {
		return err
	}
	fmt.Println(languages)
	return nil
}
