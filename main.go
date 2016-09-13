package main

import (
	"flag"
	"os"

	"github.com/intern-net/t2s/voicetext"
)

var (
	apiKey  string
	text    string
	speaker string
	out     string
)

func init() {
	flag.StringVar(&apiKey, "apikey", "", "apiKey for VoiceText API")
	flag.StringVar(&text, "text", "", "text for VoiceText API")
	flag.StringVar(&speaker, "speaker", "hikari", "speaker for VoiceText API")
	flag.StringVar(&out, "out", "", "output file name")
}

func main() {
	flag.Parse()

	dst, err := os.Create(out)
	if err != nil {
		panic(err)
	}
	err = voicetext.Text2Speech(dst, apiKey, text, speaker)
	if err != nil {
		panic(err)
	}
}
