package add

import (
	"regexp"

	"github.com/taKana671/bookmark/src/utils/bookmark"
	"github.com/taKana671/bookmark/src/utils/csv_handler"
	"github.com/taKana671/bookmark/src/utils/web"

	"github.com/spf13/cobra"
)

var (
	url      string
	category string
)

func NewAddCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add",
		Short: "Bookmark a website.",
		Long:  `Bookmark a site by saving its information in a csv file, which contains
		        the datetime when the site was bookmarked, the category, the site title,
				and the URL.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			err := run(cmd, args)
			return err
		},
	}
	cmd.Flags().StringVarP(&url, "url", "U", "", "site URL")
	cmd.Flags().StringVarP(&category, "category", "C", "all", "site category")
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
			cmd.Printf("already bookmarked on %s: %s", b.Datetime, url)
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

	cmd.Printf("bookmarked on %s: %s, %s, %s", b.Datetime, category, title, url)
	return nil

}