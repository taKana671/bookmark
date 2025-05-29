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
	data := make([][]string, 0)
	
	if !csv_handler.IsExists() {
		data = append(data, bookmark.Tags())
	} else {
		b, err := csv_handler.FindDuplication(url)

		if err != nil {
			return err
		}

		if b != nil {
			cmd.Printf("already bookmarked on %s: %s", b.Date, url)
			return nil
		}
	}

	title, err := web.GetTitle(url)
	
	if err != nil {
		return err
	}
	
	re := regexp.MustCompile(`\r?\n`)
	title = re.ReplaceAllString(title, "")
	b := bookmark.New(category, title, url)
	data = append(data, b.ToData())

	if err := csv_handler.Write(data); err != nil {
		return err
	}

	cmd.Printf("bookmarked on %s: %s, %s, %s", b.Date, category, title, url)
	return nil

}