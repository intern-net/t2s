package favclip

import (
	"fmt"
	"net/url"
	"time"
)

// ArticleGenreList represents article's genres
type ArticleGenreList struct {
	List []*ArticleGenre `json:"list"`
}

// ArticleGenre represents article's genre
type ArticleGenre struct {
	ID                 string     `json:"id"`
	Name               string     `json:"name"`
	Description        string     `json:"description,omitempty"`
	FeaturedArticlesID int64      `json:"featuredArticlesID,omitempty,string"`
	TagText            string     `json:"tagText"`
	CorePortalIDs      []string   `json:"corePortalIDs,omitempty"`
	Limits             []int      `json:"limits,omitempty"`
	IsPrivate          bool       `json:"isPrivate,omitempty"`
	CreatedAt          time.Time  `json:"createdAt,omitempty"`
	UpdatedAt          time.Time  `json:"updatedAt,omitempty"`
	Articles           []*Article `json:"articles,omitempty"`
}

// Article represents article
type Article struct {
	ID                    int64     `json:"id,string"`
	CorePortalID          string    `json:"corePortalID"`
	Title                 string    `json:"title"`
	IntroBody             string    `json:"introBody,omitempty"`
	IntroBodyHTML         string    `json:"introBodyHTML,omitempty"`
	Body                  string    `json:"body,omitempty"`
	PlainText             string    `json:"plainText,omitempty"`
	BodyHTML              string    `json:"bodyHTML,omitempty"`
	BodyContainsImageURLs []string  `json:"bodyContainsImageURLs,omitempty"`
	Status                string    `json:"status"`
	ThumbnailURL          string    `json:"thumbnailURL,omitempty"`
	EyeCatchURL           string    `json:"eyeCatchURL,omitempty"`
	PublishAt             time.Time `json:"publishAt,omitempty"`
	AutoPublishAt         time.Time `json:"autoPublishAt,omitempty"`
	AutoUnpublishAt       time.Time `json:"autoUnpublishAt,omitempty"`
	CreatedAt             time.Time `json:"createdAt,omitempty"`
	UpdatedAt             time.Time `json:"updatedAt,omitempty"`
	CanonicalURL          string    `json:"canonicalURL,omitempty"`
	MediaName             string    `json:"mediaName,omitempty"`
}

// ArticleURL returns permanent link URL
func (art *Article) ArticleURL() string {
	articleURL := &url.URL{}
	articleURL.Scheme = "https"
	articleURL.Host = fmt.Sprintf("%s.%s", art.CorePortalID, "favclip.com")
	articleURL.Path = fmt.Sprintf("/article/detail/%d", art.ID)
	return articleURL.String()
}
