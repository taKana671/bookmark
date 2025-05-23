package web

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"github.com/spf13/cobra"
)

func GetTitle(cmd *cobra.Command, url string) (string, error) {
	var title string
	resp, err := http.Get(url)

	if err != nil {
		log.Printf("failed to get html: %s", err)
		return title, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Fatalf("failed to fetch data: %d %s", resp.StatusCode, resp.Status)
		return title, fmt.Errorf("HTTP statue code: %s", resp.Status)

	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)

	if err != nil {
		log.Printf("failed to load html: %s", err)
		return title, err
	}

	title = doc.Find("title").Text()
	cmd.Println(title)
	return title, nil

}