package open

import (
	"runtime"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/taKana671/bookmark/src/utils/csv_handler"
	"github.com/taKana671/bookmark/src/utils/web"
)

var no string

func NewOpenCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "open",
		Short: "Open the bookmarked website.",
		Long:  "Locate the bookmark in the CSV file and open the website.",
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

	openCmd, err := web.NewOpenCommand(runtime.GOOS, b.Url)

	if err != nil {
		return err
	}

	if err := openCmd.Execute(); err != nil {
		return err
	}
	
	cmd.Printf("open: %s", b.Url)
	return nil
}