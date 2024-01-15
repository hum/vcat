package types

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

// Transcript is the underlying structure for processing the raw data of the transcript.
// It is used both as a serializing entity and as an output entity.
type Transcript struct {
	XMLName xml.Name `xml:"transcript" json:"-"`
	Text    []struct {
		XMLName  xml.Name `xml:"text" json:"-"`
		Start    string   `xml:"start,attr" json:"start"`  // Start time of the text
		Duration string   `xml:"dur,attr" json:"duration"` // Approximate duration of the speech in `text`
		Text     string   `xml:",innerxml"`                // The text being said in the current time bucket
	} `xml:"text" json:"data"`
}

// AvailableLanguage holds the available translation data of transcripts provided by YouTube
type AvailableLanguage struct {
	// Full name of the language translation, e.g. "English"
	Name string `json:"name"`

	// Code representation of the language's name, e.g. "en"
	Code string `json:"code"`
}
