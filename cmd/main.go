package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/hum/vcat"
)

var (
	videoURL string
)

func main() {
	flag.StringVar(&videoURL, "url", "", "url to the video to get transcription from")
	flag.StringVar(&videoURL, "u", "", "url to the video to get transcription from")
	flag.Parse()

	if videoURL == "" {
		flag.Usage()
		os.Exit(1)
	}

	captions, err := vcat.GetCaptions(videoURL)
	if err != nil {
		panic(err)
	}
	transcript, err := vcat.GetTranscript(captions.PlayerCaptionsTracklistRenderer.CaptionTracks[0].BaseUrl)
	if err != nil {
		panic(err)
	}

	fmt.Println(vcat.StringIdentStruct(transcript))
}
