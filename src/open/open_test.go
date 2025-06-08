package open

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/taKana671/bookmark/src/root"
	"github.com/taKana671/bookmark/src/utils/csv_handler"
)

func TestOpen(t *testing.T) {
	data := [][]string {
		{"datetime", "category", "title", "url"},
		{"2025-06-01 16:59:54", "golang", "Effective Go", "https://test1"},
		{"2025-06-01 16:59:55", "python", "Effective python", "https://test2"},
		{"2025-06-01 16:59:54", "google", "google", "https://www.google.co.jp"},
	}

	t.Run("successfully open site", func(t *testing.T) {
		defer rewritePath("./test_bookmarks.csv")()

		writeTestData(data, t)
		expect := "open: https://www.google.co.jp"

		tests := []struct {
			args []string
			}{
				{[]string{"open", "-N", "3"}},
				{[]string{"open", "--no", "3"}},
			}
			
			for _, tt := range tests {
				cmd := makeCmd(tt.args)
				b := bytes.NewBufferString("")
				cmd.SetOut(b)
				err := cmd.Execute()
				assert.NoError(t, err)
				
			out, err := io.ReadAll(b)
			if err != nil {
				t.Fatal(err)
			}
			
			assert.Equal(t, expect, string(out))
		}
	})

	t.Run("index out of range", func(t *testing.T) {
		defer rewritePath("./test_bookmarks.csv")()

		writeTestData(data, t)
		args := []string{"open", "-N", "5"}
		cmd := makeCmd(args)
		b := bytes.NewBufferString("")
		cmd.SetOut(b)
		err := cmd.Execute()

		expectErr := fmt.Errorf("index out of range; the number of bookmarks: %d", len(data) - 1)
		assert.Equal(t, expectErr, err)
	})

	t.Run("failed convert to int", func(t *testing.T) {
		defer rewritePath("./test_bookmarks.csv")()

		writeTestData(data, t)
		args := []string{"open", "-N", "a"}
		cmd := makeCmd(args)
		b := bytes.NewBufferString("")
		cmd.SetOut(b)
		err := cmd.Execute()
		assert.ErrorContains(t, err, `parsing "a": invalid syntax`)
	})

	t.Run("failed open file", func(t *testing.T) {
		defer rewritePath("./test_bookmarks.csv")()

		args := []string{"open", "-N", "1"}
		cmd := makeCmd(args)
		b := bytes.NewBufferString("")
		cmd.SetOut(b)
		err := cmd.Execute()
		assert.ErrorContains(t, err, "The system cannot find the file specified")
	})

}


func rewritePath(f string) func() {
	org := csv_handler.Path
	csv_handler.Path = f

	return func() {
		csv_handler.Path = org
		os.Remove(f)
	}
}

func makeCmd(args []string) *cobra.Command {
	cmd := root.NewRootCmd()
	cmd.SetArgs(args)
	searchCmd := NewOpenCmd()
	cmd.AddCommand(searchCmd)
	return cmd
}

func writeTestData(data [][]string, t *testing.T) {
	f, err := os.OpenFile(csv_handler.Path, os.O_RDWR|os.O_CREATE, 0666)
	
	if err != nil {
		t.Fatal(err)
	}
	
	defer f.Close()

	w := csv.NewWriter(f)
	w.WriteAll(data)
}