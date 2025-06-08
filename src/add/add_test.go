package add

import (
	"bytes"
	"encoding/csv"
	"io"
	"os"
	"strings"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/taKana671/bookmark/src/root"
	"github.com/taKana671/bookmark/src/utils/csv_handler"
)

func TestAdd(t *testing.T) {
	data := [][]string {
		{"datetime", "category", "title", "url"},
		{"2025-06-01 16:59:54", "golang", "Effective Go", "https://test1"},
		{"2025-06-01 16:59:55", "python", "Effective python", "https://test2"},
		{"2025-06-01 16:59:54", "golang", "start golang", "https://test3"},
	}

	t.Run("successfully bookmarked", func(t *testing.T) {
		defer rewritePath("./test_bookmarks.csv")()

		tests := []struct {
			args           []string
			expectCategory string
			expectTitle    string
			expectUrl      string

		}{
			{[]string{"add", "-U", "https://www.google.co.jp"}, "all", "Google", "https://www.google.co.jp"},
			{[]string{"add", "--category", "news", "--url", "https://www.google.co.jp"}, "news", "Google", "https://www.google.co.jp"},
			{[]string{"add", "-C", "news", "--url", "https://www.google.co.jp"}, "news", "Google", "https://www.google.co.jp"},
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

			result := string(out)
			s := strings.Split(strings.Split(result, ": ")[1], ", ")
			assert.Equal(t, tt.expectCategory, s[0])
			assert.Equal(t, tt.expectTitle, s[1])
			assert.Equal(t, tt.expectUrl, s[2])
			os.Remove(csv_handler.Path)
		}
	})
	
	t.Run("failed get title", func(t *testing.T) {
		defer rewritePath("./test_bookmarks.csv")()

		writeTestData(data, t)
		args := []string{"add", "-C", "golang", "-U", "htttps://test4"}
		cmd := makeCmd(args)
		b := bytes.NewBufferString("")
		cmd.SetOut(b)
		err := cmd.Execute()
		assert.ErrorContains(t, err, "unsupported protocol scheme")

		// file is not updated.
		rows := getFileContents(t)
		assert.Equal(t, data, rows)
	})

	t.Run("found duplication", func(t *testing.T) {
		defer rewritePath("./test_bookmarks.csv")()

		writeTestData(data, t)
		args := []string{"add", "-C", "golang", "-U", "https://test1"}
		cmd := makeCmd(args)
		b := bytes.NewBufferString("")
		cmd.SetOut(b)
		err := cmd.Execute()
		assert.NoError(t, err)

		out, err := io.ReadAll(b)
		if err != nil {
			t.Fatal(err)
		}

		expect := "already bookmarked on 2025-06-01 16:59:54: https://test1"
		assert.Equal(t, expect, string(out))
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
	searchCmd := NewAddCmd()
	cmd.AddCommand(searchCmd)
	return cmd
}


func getFileContents(t *testing.T)[][]string {
	f, err := os.Open(csv_handler.Path)

	if err != nil {
		t.Fatal(err)
	}

	defer f.Close()

	r := csv.NewReader(f)
	rows, _ := r.ReadAll()
	return rows
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