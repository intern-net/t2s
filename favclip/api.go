package favclip

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const endpoint = "https://ucon.favclip.com/api"

// NewRequest creates new request.
func NewRequest(path string, vs url.Values) (*http.Request, error) {
	u := buildURL(path, vs)
	return http.NewRequest("GET", u.String(), nil)
}

func buildURL(path string, vs url.Values) *url.URL {
	u, err := url.Parse(endpoint)
	if err != nil {
		// Must parse const endpoint
		panic(err)
	}
	u.Path += path
	if vs != nil {
		u.RawQuery = vs.Encode()
	}
	return u
}

// FetchArticleGenreList fetches article's genre list
func FetchArticleGenreList() (*ArticleGenreList, error) {
	req, err := NewRequest("/articlegenre", nil)
	if err != nil {
		return nil, err
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var data ArticleGenreList
	json.Unmarshal(body, &data)
	return &data, nil
}

// FetchArticleGenre fetches article's genre
func FetchArticleGenre(id string) (*ArticleGenre, error) {
	req, err := NewRequest(fmt.Sprintf("/articlegenre/%s", id), nil)
	if err != nil {
		return nil, err
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var data ArticleGenre
	json.Unmarshal(body, &data)
	return &data, nil
}
