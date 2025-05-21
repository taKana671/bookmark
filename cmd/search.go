/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/jszwec/csvutil"
	"github.com/spf13/cobra"
)

var (
	keyword  string
)


type Bookmark struct {
	No       int    `csv:"no"`
	Category string `csv:"category"`
	Title    string `csv:"title"`
	Url      string `csv:"url"`
}

func (b *Bookmark) CheckCategory () bool {
	if len(category) == 0 {
		return true
	}

	if b.Category == category {
		return true
	}
	return false
}


func (b *Bookmark) CheckKeyword () bool {
	if len(keyword) == 0 {
		return true
	}

	if strings.Contains(b.Title, keyword) {
		return true
	}
	return false
} 


// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("search called")

		file, err := os.Open(PATH)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		
		r := csv.NewReader(file)
		rows, err := r.ReadAll()

		if err != nil {
			log.Fatal(err)
		}

		b := &bytes.Buffer{}
		err = csv.NewWriter(b).WriteAll(rows)
		if err != nil {
			log.Fatal(err)
		}

		var bms []*Bookmark
		if err := csvutil.Unmarshal(b.Bytes(), &bms); err != nil {
			log.Fatal(err)
		}

		for _, b := range bms {
			if b.CheckCategory() && b.CheckKeyword() {
				fmt.Println(b.No, b.Category, b.Title, b.Url)
			}
		}

	},
}

func init() {
	rootCmd.AddCommand(searchCmd)
	
	searchCmd.Flags().StringVarP(&keyword, "keyword", "K", "", "keyword")
	searchCmd.Flags().StringVarP(&category, "category", "C", "", "category")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// searchCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// searchCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
