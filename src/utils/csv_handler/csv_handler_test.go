package csv_handler

import (
	"encoding/csv"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/taKana671/Bookmark/src/utils/bookmark"
)


func TestIsExists(t *testing.T) {
	t.Run("file not exist", func(t *testing.T) {
		defer rewritePath("./test_bookmarks.csv")()
		result := IsExists()
		assert.False(t, result)
	})

	t.Run("file exists", func(t *testing.T) {
		defer rewritePath("./test_bookmarks.csv")()
		f, _ := os.Create(Path)
		f.Close()
		result := IsExists()
		assert.True(t, result)
	})

}

func TestWrite(t *testing.T) {
	t.Run("failed open file", func(t *testing.T){
		defer rewritePath("./test/test_bookmarks.csv")()
		data := [][]string {
			{"datetime", "category", "title", "url"},
		}
		result := Write(data)
		assert.Error(t, result)
	})

	t.Run("successfully write", func(t *testing.T){
		defer rewritePath("./test_bookmarks.csv")()
		data := [][]string {
			{"datetime", "category", "title", "url"},
			{"2025-06-01 16:58:54", "python", "Effective Python", "https://test"},
		}
		result := Write(data)
		assert.NoError(t, result)
		
		rows := getFileContents()
		assert.Equal(t, rows, data)
	})

	t.Run("successfully append", func(t *testing.T){
		defer rewritePath("./test_bookmarks.csv")()

		data := [][]string {
			{"datetime", "category", "title", "url"},
			{"2025-06-01 16:58:54", "python", "Effective Python", "https://test"},
		}

		f, _ := os.OpenFile(Path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		w := csv.NewWriter(f)
		w.WriteAll(data)
		f.Close()

		data2 := [][]string {
			{"2025-06-01 16:59:55", "golang", "Effective go", "https://test"},
		}
		
		result := Write(data2)
		assert.NoError(t, result)
		
		rows := getFileContents()
		data = append(data, data2[0])
		assert.Equal(t, rows, data)
	})
}

func TestRead(t *testing.T) {
	t.Run("failed open file", func(t *testing.T){
		defer rewritePath("./test_bookmarks.csv")()

		result, err := Read()
		assert.Empty(t, result)
		assert.Error(t, err)
	})

	t.Run("failed read all", func(t *testing.T){
		defer rewritePath("./test_bookmarks.csv")()

		data := [][]string {
			{"datetime", "category", "title", "url"},
			{"2025-06-01 16:58:54", "Effective Python", "https://test"},
		}
	
		f, _ := os.OpenFile(Path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		w := csv.NewWriter(f)
		w.WriteAll(data)
		f.Close()

		bs, err := Read()
		assert.Empty(t, bs)
		assert.Error(t, err)
	})

	t.Run("successfully read", func(t *testing.T){
		defer rewritePath("./test_bookmarks.csv")()

		data := [][]string {
			{"datetime", "category", "title", "url"},
			{"2025-06-01 16:58:54", "python", "Effective Python", "https://test"},
		}
	
		f, _ := os.OpenFile(Path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		w := csv.NewWriter(f)
		w.WriteAll(data)
		f.Close()

		expect := &bookmark.Bookmarks{
			List: []*bookmark.Bookmark{
				{
					Datetime: "2025-06-01 16:58:54",
					Category: "python",
					Title: "Effective Python",
					Url: "https://test",
				},
			},
		}
		bs, err := Read()
		assert.Equal(t, bs, expect)
		assert.NoError(t, err)
	})
}

func TestFindDuplication(t *testing.T) {
	data := [][]string {
		{"datetime", "category", "title", "url"},
		{"2025-06-01 16:58:54", "python", "Effective Python", "https://test1"},
		{"2025-06-01 16:59:54", "golang", "Effective Go", "https://test2"},
	}

	t.Run("failed open file", func(t *testing.T){
		defer rewritePath("./test_bookmarks.csv")()

		result, err := FindDuplication("https://test")
		assert.Empty(t, result)
		assert.Error(t, err)
	})

	t.Run("Found Duplication", func(t *testing.T){
		defer rewritePath("./test_bookmarks.csv")()

		f, _ := os.OpenFile(Path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		w := csv.NewWriter(f)
		w.WriteAll(data)
		f.Close()

		expect := &bookmark.Bookmark{
			Datetime: "2025-06-01 16:58:54",
			Category: "python",
			Title: "Effective Python",
			Url: "https://test1",
		}
		b, err := FindDuplication("https://test1")
		assert.Equal(t, b, expect)
		assert.NoError(t, err)
	})

	t.Run("Found no Duplication", func(t *testing.T){
		defer rewritePath("./test_bookmarks.csv")()

		f, _ := os.OpenFile(Path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		w := csv.NewWriter(f)
		w.WriteAll(data)
		f.Close()

		b, err := FindDuplication("https://test3")
		assert.Empty(t, b)
		assert.NoError(t, err)
	})

}

func TestDelete(t *testing.T) {
	bs := &bookmark.Bookmarks{
		List: []*bookmark.Bookmark{
			{
				Datetime: "2025-06-01 16:58:54",
				Category: "python",
				Title: "Effective Python",
				Url: "https://test1",
			},
			{
				Datetime: "2025-06-01 16:59:54",
				Category: "golang",
				Title: "Effective Go",
				Url: "https://test2",
			},
		},
	}

	t.Run("failed open file", func(t *testing.T){
		defer rewritePath("./test/test_bookmarks.csv")()

		li := []*bookmark.Bookmark{}
		li = append(li, bs.List...)
		tt := &bookmark.Bookmarks{List: li}

		err := Delete(tt, 0)
		assert.Error(t, err)
	})

	t.Run("successfully delete", func(t *testing.T){
		defer rewritePath("./test_bookmarks.csv")()

		li := []*bookmark.Bookmark{}
		li = append(li, bs.List...)
		tt := &bookmark.Bookmarks{List: li}

		f, _ := os.OpenFile(Path, os.O_RDWR|os.O_CREATE, 0666)
		w := csv.NewWriter(f)
		data := tt.ToData()
		w.WriteAll(data)
		f.Close()

		expect := [][]string {
			{"datetime", "category", "title", "url"},
			{"2025-06-01 16:59:54", "golang", "Effective Go", "https://test2"},
		}

		err := Delete(tt, 0)
		assert.NoError(t, err)
		rows := getFileContents()
		assert.Equal(t, expect, rows)

	})
}

func rewritePath(f string) func() {
	org := Path
	Path = f

	return func() {
		Path = org
		os.Remove(f)
	}
}

func getFileContents()[][]string {
	f, _ := os.Open(Path)
	defer f.Close()

	r := csv.NewReader(f)
	rows, _ := r.ReadAll()
	return rows
}