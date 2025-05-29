package delete

import (
	"strconv"

	"github.com/spf13/cobra"
	"github.com/taKana671/Bookmark/src/utils/csv_handler"
)


var no string


func NewDeleteCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete",
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
	cmd.Flags().StringVarP(&no, "no", "N", "", "bookmark number")
	cmd.MarkFlagRequired("no")
	return cmd
}

func run(cmd *cobra.Command, args []string) error {
	idx, err := strconv.Atoi(no)

	if err != nil {
		return err
	}

	bs, err := csv_handler.Read()

	if err != nil {
		return err
	}

	idx -= 1
	b, err := bs.GetElement(idx)

	if err != nil {
		return err
	}

	if err := csv_handler.Delete(bs, idx); err != nil {
		return err
	}
	
	cmd.Printf("deleted a bookmark: %s, %s", b.Title, b.Url)
	return nil
}
