package vcat

import (
	"io"
	"net/http"
	"time"
)

var (
	httpclient *http.Client = &http.Client{Timeout: 60 * time.Second}
)

func GetVideo(url string) (*Video, error) {
	return getVideoWithTranscript(url, "en")
}

func GetVideoWithTranscript(url string, language string) (*Video, error) {
	return getVideoWithTranscript(url, language)
}

func GetVideoTranscript(url string, language string) ([]TextChunk, error) {
	v, err := getVideoWithTranscript(url, language)
	if err != nil {
		return nil, err
	}
	return v.Transcript, nil
}

func GetVideoMetadata(url string) (*VideoMetadata, error) {
	v, err := getVideoWithTranscript(url, "en")
	if err != nil {
		return nil, err
	}
	return v.Metadata, nil
}

func getVideoWithTranscript(url string, language string) (*Video, error) {
	captionsBody, err := do(httpclient, url)
	if err != nil {
		return nil, err
	}

	v, err := getRawVideoDetailFromInitialHttpResponse(captionsBody)
	if err != nil {
		return nil, err
	}

	var transcriptUrl = v.captions.PlayerCaptionsTracklistRenderer.CaptionTracks[0].BaseUrl

	// The default language is "en", it does not make sense to specify it twice.
	if language != "en" {
		transcriptUrl += "&tlang=" + language
	}

	transcriptBody, err := do(httpclient, transcriptUrl)
	if err != nil {
		return nil, err
	}
	transcript, err := getTranscriptFromXMLResponse(transcriptBody)
	if err != nil {
		return nil, err
	}

	var result = &Video{
		Metadata: v.metadata,
	}

	for _, text := range transcript.Text {
		result.Transcript = append(result.Transcript, TextChunk{
			Start:    text.Start,
			End:      text.End,
			Duration: text.Duration,
			Text:     text.Text,
		})
	}
	return result, nil
}

func GetAvailableCaptionLanguages(url string) ([]AvailableLanguage, error) {
	body, err := do(httpclient, url)
	if err != nil {
		return nil, err
	}

	v, err := getRawVideoDetailFromInitialHttpResponse(body)
	if err != nil {
		return nil, err
	}

	var languages = make([]AvailableLanguage, 0, len(v.captions.PlayerCaptionsTracklistRenderer.TranslationLanguages))
	for _, lang := range v.captions.PlayerCaptionsTracklistRenderer.TranslationLanguages {
		languages = append(languages, AvailableLanguage{
			Name: lang.LanguageName.SimpleText,
			Code: lang.LanguageCode,
		})
	}
	return languages, nil
}

func do(httpclient *http.Client, url string) ([]byte, error) {
	r, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	response, err := httpclient.Do(r)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
