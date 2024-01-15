# vcat

Vcat is a CLI tool to save YouTube video transcriptions without the need of an API key. Just like [cat(1)](https://man7.org/linux/man-pages/man1/cat.1.html), but for YouTube videos.
The implementation does not handle many edgecases right now.

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

List available transcriptions:

```bash
> vcat -l -u "https://www.youtube.com/watch?v=VRsbX16JAzY"
# [{Afrikaans af} {Akan ak} {Albanian sq} {Amharic am} {Arabic ar} {Armenian hy}
# {Assamese as} {Aymara ay} {Azerbaijani az} {Bangla bn} {Basque eu} {Belarusian be}
# {Bhojpuri bho} {Bosnian bs} {Bulgarian bg} {Burmese my} {Catalan ca} {Cebuano ceb}
# {Chinese (Simplified) zh-Hans} {Chinese (Traditional) zh-Hant} {Corsican co} {Croatian hr}
# {Czech cs} {Danish da} {Divehi dv} {Dutch nl} {English en} {Esperanto eo} {Estonian et}
# {Ewe ee} {Filipino fil} {Finnish fi} {French fr} {Galician gl} {Ganda lg} {Georgian ka}
# {German de} {Greek el} {Guarani gn} {Gujarati gu} {Haitian Creole ht} {Hausa ha}
# {Hawaiian haw} {Hebrew iw} {Hindi hi} {Hmong hmn} {Hungarian hu} {Icelandic is}
# {Igbo ig} {Indonesian id} {Irish ga} {Italian it} {Japanese ja} {Javanese jv} {Kannada kn}
# ...
# ]
```

Specify a different transcription language:

```bash
> vcat --language="cs" -u "https://www.youtube.com/watch?v=VRsbX16JAzY"
```

# Install

### From source

```bash
> git clone git@github.com:hum/vcat.git
> cd vcat
> make build
> ./bin/vcat -u [URL]
```

### With Go Install

```bash
> go install github.com/hum/vcat@latest
> vcat -u [URL]
```
