# vcat

Vcat helps you save YouTube video transcriptions without the need of an API key. Just like [cat(1)](https://man7.org/linux/man-pages/man1/cat.1.html), but for YouTube videos.

## Usage

#### Printing out the transcription JSON to stdout.

```bash
> vcat -u "url" --pretty
# {
#   "data": [
#     {
#       "start": "00:00:01",
#       "end": "00:00:04",
#       "duration": 3.78,
#       "text": "more recent thoughts on crypto after"
#     },
#     {
#       "start": "00:00:03",
#       "end": "00:00:07",
#       "duration": "4.981",
#       "text": "Banks not working not really"
#     },
#     ...
#}
```

#### List available transcriptions:

```bash
> vcat -l -u "url"
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

#### Specify a different transcription language:

```bash
> vcat --language="cs" -u "url"
```

#### Save to file:

```bash
> vcat -o tmp/file.json -u "url"
```

#### Work with a CSV:

```bash
> vcat -u "url" --format csv
# start,end,duration,text
# 00:00:01,00:00:04,3.78,more recent thoughts on crypto after
# 00:00:03,00:00:07,4.98,Banks not working not really
```

# Install

### From source

```bash
> git clone git@github.com:hum/vcat.git
> cd vcat
> go build -o vcat cmd/vcat/main.go
> ./vcat
```

### With Go Install

```bash
> go install github.com/hum/vcat/cmd/vcat@latest
> vcat
```
