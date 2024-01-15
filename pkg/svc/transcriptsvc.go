package svc

import (
	"io"
	"net/http"
	"time"

	"github.com/hum/vcat/pkg/svc/transformer"
	"github.com/hum/vcat/pkg/types"
)

type TranscriptSvc struct {
	httpclient *http.Client
}

func NewTranscriptSvc() *TranscriptSvc {
	httpclient := &http.Client{
		Timeout: 60 * time.Second,
	}
	return &TranscriptSvc{
		httpclient: httpclient,
	}
}

func (svc *TranscriptSvc) GetLanguageCodes(url string) ([]types.AvailableLanguage, error) {
	body, err := svc.do(url)
	if err != nil {
		return nil, err
	}

	c, err := transformer.GetCaptionsFromInitialHttpResponse(body)
	if err != nil {
		return nil, err
	}

	var languages = make([]types.AvailableLanguage, 0, len(c.PlayerCaptionsTracklistRenderer.TranslationLanguages))
	for _, lang := range c.PlayerCaptionsTracklistRenderer.TranslationLanguages {
		languages = append(languages, types.AvailableLanguage{
			Name: lang.LanguageName.SimpleText,
			Code: lang.LanguageCode,
		})
	}
	return languages, nil
}

func (svc *TranscriptSvc) GetTranscript(url string, language string) (*types.Transcript, error) {
	captionsBody, err := svc.do(url)
	if err != nil {
		return nil, err
	}

	c, err := transformer.GetCaptionsFromInitialHttpResponse(captionsBody)
	if err != nil {
		return nil, err
	}

	var transcriptUrl = c.PlayerCaptionsTracklistRenderer.CaptionTracks[0].BaseUrl

	// The default language is "en", it does not make sense to specify it twice.
	if language != "en" {
		transcriptUrl += "&tlang=" + language
	}

	transcriptBody, err := svc.do(transcriptUrl)
	if err != nil {
		return nil, err
	}
	return transformer.GetTranscriptFromXMLResponse(transcriptBody)
}

func (svc *TranscriptSvc) do(url string) ([]byte, error) {
	r, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	response, err := svc.httpclient.Do(r)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
