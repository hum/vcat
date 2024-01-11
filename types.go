package vcat

import "encoding/xml"

// Captions represent the main structure
// which holds all of the necessary data to retrieve the actual captions
type Captions struct {
	PlayerCaptionsTracklistRenderer struct {
		CaptionTracks []struct {
			BaseUrl string `json:"baseUrl"`
			Name    struct {
				SimpleText string `json:"simpleText"`
			} `json:"name"`
			LanguageCode   string `json:"languageCode"`
			Kind           string `json:"asr"`
			IsTranslatable bool   `json:"isTranslatable"`
		} `json:"captionTracks"`

		TranslationLanguages []struct {
			LanguageCode string
			LanguageName struct {
				SimpleText string `json:"simpleText"`
			} `json:"languageName"`
		} `json:"translationLanguages"`
	} `json:"playerCaptionsTracklistRenderer"`
}

type Transcript struct {
	XMLName xml.Name `xml:"transcript" json:"-"`
	Text    []struct {
		XMLName  xml.Name `xml:"text" json:"-"`
		Start    string   `xml:"start,attr"`
		Duration string   `xml:"dur,attr"`
		Context  string   `xml:",innerxml"`
	} `xml:"text"`
}
