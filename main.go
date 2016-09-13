package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/intern-net/t2s/favclip"
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

	// dst, err := os.Create(out)
	// if err != nil {
	// 	panic(err)
	// }
	// err = voicetext.Text2Speech(dst, apiKey, text, speaker)
	// if err != nil {
	// 	panic(err)
	// }

	genres, err := favclip.FetchArticleGenreList()
	if err != nil {
		panic(err)
	}

	for _, genre := range genres.List {
		genre, err := favclip.FetchArticleGenre(genre.ID)
		if err != nil {
			panic(err)
		}
		for _, art := range genre.Articles {
			if art.PlainText == "" {
				continue
			}
			dst, err := os.Create(fmt.Sprintf("voices/%d.wav", art.ID))
			if err != nil {
				panic(err)
			}
			err = voicetext.Text2Speech(dst, apiKey, art.PlainText, speaker)
			if err != nil {
				panic(err)
			}
		}
	}
}
