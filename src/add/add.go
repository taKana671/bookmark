package add

import (
	"bookmark/src/utils/bookmark"
	"bookmark/src/utils/csv_handler"
	"bookmark/src/utils/web"
	"regexp"

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
	cmd.Println(url, category)
	cmd.Println(args)

	title, err := web.GetTitle(cmd, url)
	re := regexp.MustCompile(`\r?\n`)
	title = re.ReplaceAllString(title, "")

	if err != nil {
		return err
	}

	bm := bookmark.New(category, title, url)
	
	if err := csv_handler.Write(bm.ToData()); err != nil {
		return err
	}

	return nil

}