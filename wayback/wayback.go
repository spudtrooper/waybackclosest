package wayback

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func FindClosestURL(originalURL string, raw bool) (string, error) {
	if ok, err := checkStatus(originalURL); err != nil {
		return "", err
	} else if ok {
		return originalURL, nil
	}

	res, err := waybackResults(originalURL)
	if err != nil {
		return "", err
	}

	if res.Payload.ArchivedSnapshots.Closest.Available {
		url := res.Payload.ArchivedSnapshots.Closest.URL
		if raw {
			url = strings.Replace(url, "/"+originalURL, "if_/"+originalURL, 1)
		}
		return url, nil
	}
	return "", nil
}

type Closest struct {
	URL       string `json:"url"`
	Available bool   `json:"available"`
	Status    string `json:"status"`
}

type ArchivedSnapshots struct {
	Closest `json:"closest"`
}

type Payload struct {
	URL               string `json:"url"`
	ArchivedSnapshots `json:"archived_snapshots"`
}

type results struct {
	JSON    []byte
	Payload Payload
}

func download(url string) ([]byte, error) {
	resp, err := requestResponse(url)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, errors.New(fmt.Sprintf("Error status: %d for %s", resp.StatusCode, url))
	}

	if resp.Body != nil {
		defer resp.Body.Close()
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func checkStatus(url string) (bool, error) {
	resp, err := requestResponse(url)
	if err != nil {
		return false, err
	}

	if resp.StatusCode == 200 {
		return true, nil
	}

	return false, nil
}

func requestResponse(url string) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	client := http.Client{Timeout: time.Second * 10}
	return client.Do(req)
}

func readResultsFromURL(url string) (*results, error) {
	body, err := download(url)
	if err != nil {
		return nil, err
	}
	p := &Payload{}
	if err := json.Unmarshal(body, p); err != nil {
		return nil, err
	}
	return &results{
		Payload: *p,
		JSON:    body,
	}, nil
}

func waybackResults(rawURL string) (*results, error) {
	u, err := url.Parse(rawURL)
	if err != nil {
		return nil, err
	}
	waybackURL := fmt.Sprintf("http://archive.org/wayback/available?url=%s", url.QueryEscape(u.Host+u.Path))
	res, err := readResultsFromURL(waybackURL)
	if err != nil {
		return nil, err
	}
	return res, nil
}
