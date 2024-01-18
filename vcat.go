package vcat

import (
	"io"
	"net/http"
	"time"
)

var (
	httpclient *http.Client = &http.Client{Timeout: 60 * time.Second}
)

func GetTranscription(url string, language string) (*Transcript, error) {
	captionsBody, err := do(httpclient, url)
	if err != nil {
		return nil, err
	}

	c, err := getCaptionsFromInitialHttpResponse(captionsBody)
	if err != nil {
		return nil, err
	}

	var transcriptUrl = c.PlayerCaptionsTracklistRenderer.CaptionTracks[0].BaseUrl

	// The default language is "en", it does not make sense to specify it twice.
	if language != "en" {
		transcriptUrl += "&tlang=" + language
	}

	transcriptBody, err := do(httpclient, transcriptUrl)
	if err != nil {
		return nil, err
	}
	return getTranscriptFromXMLResponse(transcriptBody)
}

func GetAvailableLanguages(url string) ([]AvailableLanguage, error) {
	body, err := do(httpclient, url)
	if err != nil {
		return nil, err
	}

	c, err := getCaptionsFromInitialHttpResponse(body)
	if err != nil {
		return nil, err
	}

	var languages = make([]AvailableLanguage, 0, len(c.PlayerCaptionsTracklistRenderer.TranslationLanguages))
	for _, lang := range c.PlayerCaptionsTracklistRenderer.TranslationLanguages {
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
