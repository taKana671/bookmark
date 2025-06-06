package search

import (
	"github.com/taKana671/Bookmark/src/utils/csv_handler"

	"github.com/spf13/cobra"
)

var (
	category string
	keyword  string
)

func NewSearchCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "search",
		Short: "A brief description of your command",
		Long: `A longer description that spans multiple lines and likely contains examples
				and usage of using your command. For example:

				Cobra is a CLI library for Go that empowers applications.
				This application is a tool to generate the needed files
				to quickly create a Cobra application.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			err := run(cmd, args)
			return err
		},
	}
	cmd.Flags().StringVarP(&keyword, "keyword", "K", "", "keyword")
	cmd.Flags().StringVarP(&category, "category", "C", "", "category")
	return cmd
}

func run(cmd *cobra.Command, args []string) error {
	bs, err := csv_handler.Read()

	if err != nil {
		return err
	}

	if len(bs.List) == 0 {
		cmd.Printf("the number of bookmarks: %d", len(bs.List))
		return nil
	}

	cnt := 0

	for _, b := range bs.List {
		if b.CheckCategory(category) && b.CheckKeyword(keyword) {
			cnt++
			cmd.Println(cnt, b.ToData())
		}
	}

	if cnt == 0 {
		cmd.Printf("bookmarks not found; category: %s; keyword: %s", category, keyword)
	}

	return nil
}