package favclip

import (
	"net/url"
	"testing"
)

func TestBuildURL(t *testing.T) {
	{
		u := buildURL("/path/to/api", nil)
		exp := "https://ucon.favclip.com/api/path/to/api"
		if u.String() != exp {
			t.Errorf("expected: %s, but got: %s", exp, u.String())
		}
	}

	{
		vs := url.Values{}
		vs.Set("foo", "bar")
		vs.Add("tags", "xxx")
		vs.Add("tags", "yyy")
		u := buildURL("/path/to/api", vs)
		exp := "https://ucon.favclip.com/api/path/to/api?foo=bar&tags=xxx&tags=yyy"
		if u.String() != exp {
			t.Errorf("expected: %s, but got: %s", exp, u.String())
		}
	}
}

func TestFetchArticleGenreList(t *testing.T) {
	genres, _ := FetchArticleGenreList()
	exp := 10
	got := len(genres.List)
	if got != exp {
		t.Errorf("expected: %d, but got: %d", exp, got)
	}
}

func TestFetchArticleGenre(t *testing.T) {
	genre, _ := FetchArticleGenre("all")
	exp := "ALL"
	got := genre.Name
	if got != exp {
		t.Errorf("expected: %s, but got: %s", exp, got)
	}
}
