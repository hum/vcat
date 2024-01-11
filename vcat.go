package vcat

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
