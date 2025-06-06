package search

import (
	"bytes"
	"encoding/csv"
	"io"
	"os"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/taKana671/Bookmark/src/root"
	"github.com/taKana671/Bookmark/src/utils/csv_handler"
)

func TestSearch(t *testing.T) {
	data := [][]string {
		{"datetime", "category", "title", "url"},
		{"2025-06-01 16:59:54", "golang", "Effective Go", "https://test1"},
		{"2025-06-01 16:59:55", "python", "Effective python", "https://test2"},
		{"2025-06-01 16:59:54", "golang", "start golang", "https://test3"},
	}

	t.Run("search with no options", func(t *testing.T) {
		defer rewritePath("./test_bookmarks.csv")()
		
		f, _ := os.OpenFile(csv_handler.Path, os.O_RDWR|os.O_CREATE, 0666)
		w := csv.NewWriter(f)
		w.WriteAll(data)
		f.Close()

		args := []string{"search"}
		cmd := makeCmd(args)
		
		b := bytes.NewBufferString("")
		cmd.SetOut(b)
		err := cmd.Execute()
		assert.NoError(t, err)

		expect := "1 [2025-06-01 16:59:54 golang Effective Go https://test1]\n2 [2025-06-01 16:59:55 python Effective python https://test2]\n3 [2025-06-01 16:59:54 golang start golang https://test3]\n"

		out, err := io.ReadAll(b)
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, expect, string(out))
	})

	t.Run("search with options", func(t *testing.T) {
		defer rewritePath("./test_bookmarks.csv")()
		
		f, _ := os.OpenFile(csv_handler.Path, os.O_RDWR|os.O_CREATE, 0666)
		w := csv.NewWriter(f)
		w.WriteAll(data)
		f.Close()

		tests := []struct {
			args []string
			expect string
		}{
			{[]string{"search", "--category", "python"}, "1 [2025-06-01 16:59:55 python Effective python https://test2]\n"},
			{[]string{"search", "-C", "python"}, "1 [2025-06-01 16:59:55 python Effective python https://test2]\n"},
			{[]string{"search", "--keyword", "start"}, "1 [2025-06-01 16:59:54 golang start golang https://test3]\n"},
			{[]string{"search", "-K", "start"}, "1 [2025-06-01 16:59:54 golang start golang https://test3]\n"},
			{[]string{"search", "--category", "golang", "--keyword", "Effective"}, "1 [2025-06-01 16:59:54 golang Effective Go https://test1]\n"},
			{[]string{"search", "-C", "golang", "--keyword", "Effective"}, "1 [2025-06-01 16:59:54 golang Effective Go https://test1]\n"},
			{[]string{"search", "--category", "golang", "-K", "Effective"}, "1 [2025-06-01 16:59:54 golang Effective Go https://test1]\n"},
			
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
			assert.Equal(t, tt.expect, string(out))
		}
	})

	t.Run("no match", func(t *testing.T) {
		defer rewritePath("./test_bookmarks.csv")()

		f, _ := os.OpenFile(csv_handler.Path, os.O_RDWR|os.O_CREATE, 0666)
		w := csv.NewWriter(f)
		w.WriteAll(data)
		f.Close()

		args := []string{"search", "--category", "SQL"}
		cmd := makeCmd(args)
		
		b := bytes.NewBufferString("")
		cmd.SetOut(b)
		err := cmd.Execute()
		assert.NoError(t, err)

		out, err := io.ReadAll(b)
		if err != nil {
			t.Fatal(err)
		}

		expect := "bookmarks not found; category: SQL; keyword: "
		assert.Equal(t, expect, string(out))
	})

	t.Run("no bookmarks in file", func(t *testing.T) {
		defer rewritePath("./test_bookmarks.csv")()
		
		title := [][]string {
			{"datetime", "category", "title", "url"},
		}

		f, _ := os.OpenFile(csv_handler.Path, os.O_RDWR|os.O_CREATE, 0666)
		w := csv.NewWriter(f)
		w.WriteAll(title)
		f.Close()

		args := []string{"search", "--category", "SQL"}
		cmd := makeCmd(args)
		
		b := bytes.NewBufferString("")
		cmd.SetOut(b)
		err := cmd.Execute()
		assert.NoError(t, err)

		out, err := io.ReadAll(b)
		if err != nil {
			t.Fatal(err)
		}

		expect := "the number of bookmarks: 0"
		assert.Equal(t, expect, string(out))
	})

	t.Run("file open error", func(t *testing.T) {
		defer rewritePath("./test_bookmarks.csv")()

		args := []string{"search", "--category", "SQL"}
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
	searchCmd := NewSearchCmd()
	cmd.AddCommand(searchCmd)
	return cmd
}