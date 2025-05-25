package add

import (
	"regexp"

	"github.com/taKana671/Bookmark/src/utils/bookmark"
	"github.com/taKana671/Bookmark/src/utils/csv_handler"
	"github.com/taKana671/Bookmark/src/utils/web"

	"github.com/spf13/cobra"
)

var (
	url      string
	category string
)

const PATH = "bookmarks.csv"

func NewAddCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add",
		Short: "A brief description of your command",
		Long:  `A longer description that spans multiple lines and likely contains examples
				and usage of using your command. For example:
				
				Cobra is a CLI library for Go that empowers applications.
				This application is a tool to generate the needed files
				to quickly create a Cobra application.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			err := run(cmd, args)
			return err
		},
	}
	cmd.Flags().StringVarP(&url, "url", "U", "", "site URL")
	cmd.Flags().StringVarP(&category, "category", "C", "all", "site URL")
	cmd.MarkFlagRequired("url")
	return cmd
}

func run(cmd *cobra.Command, args []string) error {
	title, err := web.GetTitle(cmd, url)
	
	if err != nil {
		return err
	}
	
	re := regexp.MustCompile(`\r?\n`)
	title = re.ReplaceAllString(title, "")
	bm := bookmark.New(category, title, url)
	data := make([][]string, 0)
	
	if ! csv_handler.IsExists() {
		data = append(data, bm.Fields())
	} else {
		// check duplication
		// if th same url is found, return nil and println(already bookmarked)
	}

	data = append(data, bm.ToData())

	if err := csv_handler.Write(cmd, data); err != nil {
		return err
	}

	cmd.Printf("add data: %s, %s", category, url)
	return nil

}