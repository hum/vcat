# vcat

Vcat is a CLI tool to save YouTube video transcriptions without the need of an API key. Just like [cat(1)](https://man7.org/linux/man-pages/man1/cat.1.html), but for YouTube videos.
The implementation is very simple and does not handle any edgecases right now.

## Usage

Printing out the transcription JSON to stdout.

```bash
> ./bin/vcat -u "https://www.youtube.com/watch?v=VRsbX16JAzY"
# {
#	"Text": [
#		{
#			"Start": "1.38",
#			"Duration": "3.78",
#			"Context": "more recent thoughts on crypto after"
#		},
#		{
#			"Start": "3.179",
#			"Duration": "4.981",
#			"Context": "Banks not working not really"
#		},
#    ...
#}
```
