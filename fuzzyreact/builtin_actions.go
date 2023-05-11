package fuzzyreact

import (
	"encoding/json"
	"net/http"
	"net/url"
	"time"
)

type WikiResult struct {
	Batchcomplete string    `json:"batchcomplete"`
	Continue      any       `json:"continue"`
	Query         WikiQuery `json:"query"`
}

type WikiQuery struct {
	Searchinfo any          `json:"searchinfo"`
	Search     []WikiSearch `json:"search"`
}

type WikiSearch struct {
	Ns        int       `json:"ns"`
	Title     string    `json:"title"`
	Pageid    int       `json:"pageid"`
	Size      int       `json:"size"`
	Wordcount int       `json:"wordcount"`
	Snippet   string    `json:"snippet"`
	Timestamp time.Time `json:"timestamp"`
}

func Wikipedia(q string) (string, error) {

	u, err := url.Parse("https://en.wikipedia.org/w/api.php?action=query&list=search&format=json")
	if err != nil {
		return "", err
	}

	query := u.Query()
	query.Add("srsearch", q)
	u.RawQuery = query.Encode()

	resp, err := http.Get(u.String())
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	data := &WikiResult{}
	if err := json.NewDecoder(resp.Body).Decode(data); err != nil {
		return "", err
	}

	return data.Query.Search[0].Snippet, nil

}
