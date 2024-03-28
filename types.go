package vcat

import (
	"encoding/xml"
	"errors"
)

var (
	ErrTranscriptNotFound error = errors.New("no trancript found for given url")
	ErrCaptionsNotFound   error = errors.New("no captions found for url")
)

type Video struct {
	Metadata   *VideoMetadata `json:"metadata"`
	Transcript []TextChunk    `json:"transcript"`
}

// VideoMetadata stores information related to the video, e.g. the title, or the thumbnails
type VideoMetadata struct {
	VideoId          string   `json:"videoId"`
	Title            string   `json:"title"`
	LengthSeconds    string   `json:"lengthSeconds"`
	Keywords         []string `json:"keywords"`
	ChannelId        string   `json:"channelId"`
	ShortDescription string   `json:"shortDescription"`
	Thumbnail        struct {
		Thumbnails []struct {
			Url    string `json:"url"`
			Width  int    `json:"width"`
			Height int    `json:"height"`
		} `json:"thumbnails"`
	} `json:"thumbnail"`
	ViewCount     string `json:"viewCount"`
	Author        string `json:"author"`
	IsPrivate     bool   `json:"isPrivate"`
	IsLiveContent bool   `json:"isLiveContent"`
}

type TextChunk struct {
	Start    string  `json:"start"`    // Start time of the text
	End      string  `json:"end"`      // End time of the text
	Duration float64 `json:"duration"` // Approximate duration of the speech in `text`
	Text     string  `json:"text"`     // The text being said in the current time bucket
}

// AvailableLanguage holds the available translation data of transcripts provided by YouTube
type AvailableLanguage struct {
	// Full name of the language translation, e.g. "English"
	Name string `json:"name"`

	// Code representation of the language's name, e.g. "en"
	Code string `json:"code"`
}

// Captions represent the main structure
// which holds all of the necessary data to retrieve the actual captions
type captions struct {
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
type transcript struct {
	XMLName xml.Name `xml:"transcript" json:"-"`
	Text    []struct {
		XMLName  xml.Name `xml:"text" json:"-"`
		Start    string   `xml:"start,attr" json:"start"` // Start time of the text
		End      string   `json:"end"`
		Duration float64  `xml:"dur,attr" json:"duration"` // Approximate duration of the speech in `text`
		Text     string   `xml:",innerxml" json:"text"`    // The text being said in the current time bucket
	} `xml:"text" json:"data"`
}

type rawVideoDetail struct {
	metadata *VideoMetadata
	captions captions
}
