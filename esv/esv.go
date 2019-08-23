package esv

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
)

// Verse defines the response from the ESV api
type Verse struct {
	Query    string
	Passages []string
}

func fetchVerse(body []byte) (*Verse, error) {
	v := new(Verse)
	err := json.Unmarshal(body, &v)
	return v, err
}

// GetVerse returns a random verse of the bible
func GetVerse(search string) (string, error) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", "https://api.esv.org/v3/passage/text/?q="+search, nil)
	if err != nil {
		return "", err
	}

	req.Header.Add("Authorization", "Token "+os.Getenv("CROSSWAY_API_KEY"))

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	verse, err := fetchVerse(body)
	if err != nil {
		return "", err
	}

	return verse.Passages[0], err
}

// GetVerseHTML returns the verse HTML
func GetVerseHTML(search string) (string, error) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", "https://api.esv.org/v3/passage/html/?q="+search, nil)
	if err != nil {
		return "", err
	}

	req.Header.Add("Authorization", "Token "+os.Getenv("CROSSWAY_API_KEY"))

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	verse, err := fetchVerse(body)
	if err != nil {
		return "", err
	}

	return verse.Passages[0], err
}
