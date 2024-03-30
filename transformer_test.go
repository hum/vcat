package vcat

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTransformerReturnsRawVideoDetailShouldPass(t *testing.T) {
	tests := []struct {
		url string
	}{
		{
			url: "https://www.youtube.com/watch?v=uyMtsyzXWd4",
		},
		{
			url: "https://www.youtube.com/watch?v=709z-t7IiFw",
		},
	}

	for _, tt := range tests {
		body, err := do(httpclient, tt.url)
		require.NoError(t, err)

		output, err := getRawVideoDetailFromInitialHttpResponse(body)
		require.NoError(t, err)
		require.NotEmpty(t, output)
	}
}

func TestTransformerGetsTranscriptFromURLShouldPass(t *testing.T) {
	tests := []struct {
		url string
	}{
		{
			url: "https://www.youtube.com/watch?v=uyMtsyzXWd4",
		},
		{
			url: "https://www.youtube.com/watch?v=709z-t7IiFw",
		},
	}

	for _, tt := range tests {
		rawb, err := do(httpclient, tt.url)
		require.NoError(t, err)

		rawv, err := getRawVideoDetailFromInitialHttpResponse(rawb)
		require.NoError(t, err)

		var transcriptUrl = rawv.captions.PlayerCaptionsTracklistRenderer.CaptionTracks[0].BaseUrl
		body, err := do(httpclient, transcriptUrl)
		require.NoError(t, err)

		transcript, err := getTranscriptFromXMLResponse(body)
		require.NoError(t, err)

		require.NotEmpty(t, transcript)
		require.NotEmpty(t, transcript.Text)
	}
}
