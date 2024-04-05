package vcat

import "fmt"

// GetVideoWithLanguage returns the whole video data, including the transcript in a specified language.
// Do pass in the language code, not the name of the language. E.g. "en", not "English".
//
// Make sure the provided language is supported by asserting it is available on the content.
// You can use `vcat.GetAvailableLanguages`.
func GetVideoWithLanguage(url, languageCode string) (*Video, error) {
	return getVideo(url, languageCode)
}

// GetVideo returns the whole video data, including the transcript, in English.
//
// An alias for `vcat.GetVideoWithLanguage(url, "en")`
func GetVideo(url string) (*Video, error) {
	return GetVideoWithLanguage(url, "en")
}

// GetAvailableLanguages returns all valid transcript languages available for the specified url.
//
// The returned language names (not codes) could be translated to the language in your location.
// I.e. if the process calling this function has a Spanish IP, the names of the available langugues are going to be in Spanish.
func GetAvailableLanguages(url string) ([]AvailableLanguage, error) {
	v, err := getVideoDetail(url)
	if err != nil {
		return nil, err
	}

	var languages = make([]AvailableLanguage, 0, len(v.captions.PlayerCaptionsTracklistRenderer.TranslationLanguages))
	for _, l := range v.captions.PlayerCaptionsTracklistRenderer.TranslationLanguages {
		languages = append(languages, AvailableLanguage{
			Name: l.LanguageName.SimpleText,
			Code: l.LanguageCode,
		})
	}
	return languages, nil
}

// getVideo takes in a base URL for the video, and the language code, to return the video metadata along with the transcript.
//
// It is up to the caller to validate the passed in languageCode is supported for this video URL.
func getVideo(url, languageCode string) (*Video, error) {
	detail, err := getVideoDetail(url)
	if err != nil {
		return nil, err
	}

	var transcriptUrl = detail.captions.PlayerCaptionsTracklistRenderer.CaptionTracks[0].BaseUrl

	// Only include the language if it isn't English
	if languageCode != "en" {
		// @TODO: Should we validate the passed in languageCode is supported for this video? Otherwise we are wasting http requests.
		transcriptUrl += "&tlang=" + languageCode
	}

	chunks, err := getTranscriptFromUrl(transcriptUrl)
	if err != nil {
		return nil, err
	}

	return &Video{
		Metadata:   detail.metadata,
		Transcript: chunks,
	}, nil
}

// getVideoDetail returns the raw detail from the base video URL
func getVideoDetail(url string) (*rawVideoDetail, error) {
	body, err := do(httpclient, url)
	if err != nil {
		return nil, fmt.Errorf("could not request video detail, got err=%s", err)
	}
	return getRawVideoDetailFromInitialHttpResponse(body)
}

// getTranscriptFromUrl takes in the actual transcript URL and parses it into a slice of TranscriptTextChunks.
//
// The provided transcript is returned as-is without any special chunking. It is the raw YouTube transcript.
func getTranscriptFromUrl(transcriptUrl string) ([]TranscriptTextChunk, error) {
	body, err := do(httpclient, transcriptUrl)
	if err != nil {
		return nil, err
	}

	t, err := getTranscriptFromXMLResponse(body)
	if err != nil {
		return nil, err
	}

	var chunks = make([]TranscriptTextChunk, 0, len(t.Text))
	for _, text := range t.Text {
		chunks = append(chunks, TranscriptTextChunk{
			Start:    text.Start,
			End:      text.End,
			Duration: text.Duration,
			Text:     text.Text,
		})
	}
	return chunks, nil
}
