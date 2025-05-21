/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/PuerkitoBio/goquery"
	"github.com/spf13/cobra"
)

var (
	url      string
    category string
)

// const PATH = "C:/Users/Kanae/Desktop/go_development/Bookmarks"
const PATH = "bookmarks.csv"


// type Bookmarks struct {
// 	Roots Roots `json:"roots"`

// }

// type Roots struct {
// 	Other Other `json:"other"`
// }

// type Other struct {
// 	Children []Children `json:"children"`
	// Name string `json:"name"`
	// Id string `json:"id"`
	// Type string `json:"type"`
// }

// type Children struct {
	
// 	Children []Child `json:"children"`
// 	// Children Other `json:"children,omitempty"`
// 	// Children []Other `json:"children"`
// }

// type Children struct {
// 	Id string `json:"id"`
// 	Name string `json:"name"`
// 	Type string `json:"type"`
// 	Url string `json:"url"`
// 	Children []Children `json:"children"`
	// Children Other `json:"children,omitempty"`
	// Children []Other `json:"children"`
// }

// type Child struct {
// 	Id string `json:"id"`
// 	Name string `json:"name"`
// 	Type string `json:"type"`
// 	Url string `json:"url"`
// 	Children []*Child `json:"children,omitempty"`
// }


// type Bookmarks struct {
// 	Category string ``
// }



func getTitle() string {
	resp, err := http.Get(url)

	if err != nil {
		log.Printf("failed to get html: %s", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Fatalf("failed to fetch data: %d %s", resp.StatusCode, resp.Status)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)

	if err != nil {
		log.Printf("failed to load html: %s", err)
	}

	title := doc.Find("title").Text()
	fmt.Println(title)
	return title

}


// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
	and usage of using your command. For example:
	
	Cobra is a CLI library for Go that empowers applications.
	This application is a tool to generate the needed files
	to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(url, category)

		title := getTitle()

		data := []string{
			category,
			title,
			url,
		}
		f, err := os.OpenFile(PATH, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

		if err != nil {
			log.Printf("failed to open %s: %s", PATH, err)
		}

		defer f.Close()

		w := csv.NewWriter(f)
		if err := w.Write(data); err != nil {
			log.Printf("failed to write %s: %s", PATH, err)
		}

		w.Flush()


		// fmt.Println("add called")

		// resp, err := http.Get(site)

		// if err != nil {
		// 	log.Printf("failed to get html: %s", err)
		// }

		// defer resp.Body.Close()

		// if resp.StatusCode != 200 {
		// 	log.Fatalf("failed to fetch data: %d %s", resp.StatusCode, resp.Status)
		// }

		// doc, err := goquery.NewDocumentFromReader(resp.Body)

		// if err != nil {
		// 	log.Printf("failed to load html: %s", err)
		// }

		// title := doc.Find("title").Text()
		// fmt.Println(title)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	
	addCmd.Flags().StringVarP(&url, "url", "U", "", "site URL")
	addCmd.Flags().StringVarP(&category, "category", "C", "all", "site URL")
	// https://gucchi.blog/221017_bookmarks/

	//必須のフラグに指定
	addCmd.MarkFlagRequired("site")
}

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

