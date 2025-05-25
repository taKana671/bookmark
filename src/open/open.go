package open

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/taKana671/Bookmark/src/utils/csv_handler"
	"github.com/taKana671/Bookmark/src/utils/web"
)

var (
	no  string
	url string
)

func NewOpenCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "open",
		Short: "A brief description of your command",
		Long: `A longer description that spans multiple lines and likely contains examples
				and usage of using your command. For example:

				Cobra is a CLI library for Go that empowers applications.
				This application is a tool to generate the needed files
				to quickly create a Cobra application.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			// fmt.Println("open called")
			err := run(cmd, args)
			return err
		},
	}
	cmd.Flags().StringVarP(&no, "no", "N", "", "bookmark number")
	cmd.Flags().StringVarP(&url, "url", "U", "", "site URL")
	return cmd
}

func run(cmd *cobra.Command, args []string) error {
	idx, err := strconv.Atoi(no)

	if err != nil {
		cmd.PrintErrf("cannot convert to integer: %s", no)
		return err
	}

	bms, err := csv_handler.Read(cmd)

	if err != nil {
		cmd.PrintErrln(err)
		return err
	}

	idx -= 1
	if idx < 0 || idx >= len(bms) {
		return fmt.Errorf("outof index: %s", no)
	}

	bm := bms[idx]
	err = web.Open(cmd, bm.Url)
	
	if err != nil {
		cmd.PrintErrf("cannot open site: %s", bm.Url)
		return err
	}
	
	return nil
	
}