package delete

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"os"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/taKana671/bookmark/src/root"
	"github.com/taKana671/bookmark/src/utils/csv_handler"
)

func TestDelete(t *testing.T) {
	data := [][]string {
		{"datetime", "category", "title", "url"},
		{"2025-06-01 16:59:54", "golang", "Effective Go", "https://test1"},
		{"2025-06-01 16:59:55", "python", "Effective python", "https://test2"},
		{"2025-06-01 16:59:54", "golang", "start golang", "https://test3"},
	}

	t.Run("successfully delete", func(t *testing.T) {
		defer rewritePath("./test_bookmarks.csv")()

		tests := []struct {
			args []string
		}{
			{[]string{"delete", "-N", "3"}},
			{[]string{"delete", "--no", "3"}},
			
		}

		for _, tt := range(tests) {
			writeTestData(data, t)
			cmd := makeCmd(tt.args)
			b := bytes.NewBufferString("")
			cmd.SetOut(b)
			err := cmd.Execute()
			fmt.Println(err)

			assert.NoError(t, err)
			rows := getFileContents(t)
			assert.Equal(t, data[:3], rows)

			os.Remove(csv_handler.Path)
		}
	})


	t.Run("failed convert to int", func(t *testing.T) {
		defer rewritePath("./test_bookmarks.csv")()

		writeTestData(data, t)
		args := []string{"delete", "-N", "a"}
		cmd := makeCmd(args)
		b := bytes.NewBufferString("")
		cmd.SetOut(b)
		err := cmd.Execute()

		assert.ErrorContains(t, err, `parsing "a": invalid syntax`)
		rows := getFileContents(t)
		assert.Equal(t, data, rows)

	})

	t.Run("failed csv read", func(t *testing.T) {
		defer rewritePath(".test/test_bookmarks.csv")()

		args := []string{"delete", "-N", "1"}
		cmd := makeCmd(args)
		b := bytes.NewBufferString("")
		cmd.SetOut(b)
		err := cmd.Execute()
		
		assert.ErrorContains(t, err, "The system cannot find the path specified.")
	})

	t.Run("index out of range", func(t *testing.T) {
		defer rewritePath("./test_bookmarks.csv")()

		writeTestData(data, t)
		args := []string{"delete", "-N", "0"}
		cmd := makeCmd(args)
		b := bytes.NewBufferString("")
		cmd.SetOut(b)
		err := cmd.Execute()

		assert.ErrorContains(t, err, "index out of range; the number of bookmarks: 3")
		rows := getFileContents(t)
		assert.Equal(t, data, rows)
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
	searchCmd := NewDeleteCmd()
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