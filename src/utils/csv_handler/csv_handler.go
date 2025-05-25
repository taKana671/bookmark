package csv_handler

import (
	"bytes"
	"encoding/csv"
	"os"

	"github.com/spf13/cobra"
	"github.com/taKana671/Bookmark/src/utils/bookmark"

	"github.com/jszwec/csvutil"
)

const PATH = "bookmarks.csv"


func IsExists() bool {
	if _, err := os.Stat(PATH); err == nil {
		return true
	}

	return false
}

func Write(cmd *cobra.Command, data [][]string) error {
	f, err := os.OpenFile(PATH, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		cmd.PrintErrf("failed to open %s: %s", PATH, err)
		return err
	}

	defer f.Close()

	w := csv.NewWriter(f)

	if err := w.WriteAll(data); err != nil {
		cmd.PrintErrf("failed to write %s: %s", PATH, err)
		return err
	}
	// }	if err := w.Write(data); err != nil {
	// 	cmd.PrintErrf("failed to write %s: %s", PATH, err)
	// 	return err
	// }

	w.Flush()
	return nil
}

func Read(cmd *cobra.Command) ([]*bookmark.Bookmark, error) {
	var bms []*bookmark.Bookmark

	f, err := os.Open(PATH)
	
	if err != nil {
		cmd.PrintErrln(err)
		return bms, err
	}

	defer f.Close()
	
	r := csv.NewReader(f)
	rows, err := r.ReadAll()

	if err != nil {
		cmd.PrintErrln(err)
		return bms, err
	}

	b := &bytes.Buffer{}
	if err := csv.NewWriter(b).WriteAll(rows); err != nil {
		cmd.PrintErrln(err)
		return bms, err
	}

	if err := csvutil.Unmarshal(b.Bytes(), &bms); err != nil {
		cmd.PrintErrln(err)
		return bms, err
	}

	return bms, nil

}