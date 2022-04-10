package receive

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/PuerkitoBio/goquery"
)

// message format
type Message map[string]interface{}

func Deserialize(b []byte) (Message, error) {
	var msg Message
	buf := bytes.NewBuffer(b)
	decoder := json.NewDecoder(buf)
	err := decoder.Decode(&msg)
	return msg, err
}

func getTitles(title, lang string) []string {
	url := `https://` + lang + `.wikipedia.org/w/api.php?action=query&prop=links&pllimit=max&format=json&titles=` + url.QueryEscape(title)

	resp, err := http.Get(url)
	if err != nil {
		log.Println("get request failed", url, err)
		return nil
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("failed to read body", url, err)
		return nil
	}

	msg, err := Deserialize(body)
	if err != nil {
		log.Println("failed to deserialize body", url, err, string(body))
	}

	titles := make([]string, 0)

	query := msg["query"].(map[string]interface{})
	pages := query["pages"].(map[string]interface{})
	for _, page := range pages {
		links := page.(map[string]interface{})["links"]
		if links == nil {
			continue
		}
		for _, link := range links.([]interface{}) {
			title := link.(map[string]interface{})["title"].(string)
			titles = append(titles, title)
		}
	}

	return titles
}

func extractLinks(pageUrl string) []string {
	// parse base url
	base, err := url.Parse(pageUrl)
	if err != nil {
		fmt.Fprintln(os.Stderr, "failed to parse page url", pageUrl, err)
		return nil
	}

	// send get request and receive response from server
	resp, err := http.Get(pageUrl)
	if err != nil {
		fmt.Fprintln(os.Stderr, "get request failed", pageUrl, err)
		return nil
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		fmt.Fprintln(os.Stderr, "got non 200 status code", pageUrl, resp.StatusCode)
		return nil
	}

	// extract links
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to parse body", pageUrl, err)
		return nil
	}

	hrefs := make([]string, 0)
	doc.Find("a[href]").Each(func(index int, item *goquery.Selection) {
		href, _ := item.Attr("href")

		absolute, err := base.Parse(href)
		if err != nil {
			fmt.Fprintln(os.Stderr, "failed to parse href", href, err)
			return
		}

		href = absolute.String()

		hrefs = append(hrefs, href)
	})

	return hrefs
}
