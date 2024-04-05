package vcat

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/require"
)

var videourl = "https://www.youtube.com/watch?v=Nb2tebYAaOA" // Jim Keller interview

func TestGetVideoWithLanguageShouldPass(t *testing.T) {
	languages, err := GetAvailableLanguages(videourl)
	require.NoError(t, err)
	require.NotEmpty(t, languages)

	// Pick a random language from the available languages and fetch the transcript for the language
	lang := languages[rand.Intn(len(languages))]

	video, err := GetVideoWithLanguage(videourl, lang.Code)
	require.NoError(t, err)
	require.NotEmpty(t, video)
	require.NotEmpty(t, video.Metadata.Title)
	require.NotEmpty(t, video.Transcript)
}

func TestGetAvailableLanguagesReturnsNonZeroSliceShouldPass(t *testing.T) {
	languages, err := GetAvailableLanguages(videourl)
	require.NoError(t, err)
	require.NotEmpty(t, languages)
}
