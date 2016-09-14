package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/intern-net/t2s/favclip"
	"github.com/intern-net/t2s/say"
)

var (
	apiKey  string
	text    string
	speaker string
	out     string
)

// Voice represents voice meta
type Voice struct {
	ID      int64  `json:"id"`
	Title   string `json:"title"`
	LinkURL string `json:"linkURL"`
}

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
		// make directory
		dirname := fmt.Sprintf("says/%s", genre.ID)
		err = os.MkdirAll(dirname, 0777)
		if err != nil {
			panic(err)
		}
		for _, art := range genre.Articles {
			if art.PlainText == "" {
				continue
			}
			// create voice file
			voicefile := fmt.Sprintf("%s/%d.mp4", dirname, art.ID)
			err := say.Text2Speech(voicefile, art.PlainText, speaker)
			if err != nil {
				panic(err)
			}
			// create meta json
			metafile := fmt.Sprintf("%s/%d.json", dirname, art.ID)
			v := &Voice{
				ID:      art.ID,
				Title:   art.Title,
				LinkURL: art.ArticleURL(),
			}
			b, err := json.MarshalIndent(v, "", "  ")
			if err != nil {
				panic(err)
			}
			err = ioutil.WriteFile(metafile, b, os.ModePerm)
			if err != nil {
				panic(err)
			}
		}
	}
}
