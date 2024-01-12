package main

import (
	"flag"
	"fmt"
	"os"
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

	captions, err := GetCaptions(videoURL)
	if err != nil {
		panic(err)
	}
	transcript, err := GetTranscript(captions.PlayerCaptionsTracklistRenderer.CaptionTracks[0].BaseUrl)
	if err != nil {
		panic(err)
	}

	fmt.Println(StringIdentStruct(transcript))
}

// GetCaptions is the entrypoint to fetching video captions for a YouTube video.
//
// @TODO: Handle edge-cases.
func GetCaptions(url string) (Captions, error) {
	b, err := GetBodyAsByteSlice(url)
	if err != nil {
		return Captions{}, err
	}
	return GetCaptionsFromRawHtml(b)
}

// GetTranscript is the entrypoint for fetching the actual transcript for a YouTube video.
func GetTranscript(url string) (Transcript, error) {
	b, err := GetBodyAsByteSlice(url)
	if err != nil {
		return Transcript{}, nil
	}
	return ParseTranscriptFromXml(b)
}
