package csv_handler

import (
	"bytes"
	"encoding/csv"
	"os"
	"slices"

	"github.com/taKana671/Bookmark/src/utils/bookmark"

	"github.com/jszwec/csvutil"
)

var Path string = "bookmarks.csv"


func IsExists() bool {
	if _, err := os.Stat(Path); err == nil {
		return true
	}

	return false
}

func Write(data [][]string) error {
	f, err := os.OpenFile(Path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		return err
	}

	defer f.Close()

	w := csv.NewWriter(f)

	if err := w.WriteAll(data); err != nil {
		return err
	}
	
	// w.Flush()
	return nil
}

func Read() (*bookmark.Bookmarks, error) {
	f, err := os.Open(Path)
	
	if err != nil {
		return nil, err
	}

	defer f.Close()
	
	r := csv.NewReader(f)
	rows, err := r.ReadAll()

	if err != nil {
		return nil, err
	}

	b := &bytes.Buffer{}
	
	if err := csv.NewWriter(b).WriteAll(rows); err != nil {
		return nil, err
		
	}
	
	var list []*bookmark.Bookmark

	if err := csvutil.Unmarshal(b.Bytes(), &list); err != nil {
		return nil, err
	}

	bs := &bookmark.Bookmarks{List: list}
	return bs, nil
}

func FindDuplication(url string) (*bookmark.Bookmark, error) {
	bs, err := Read()

	if err != nil {
		return nil, err
	}

	for _, b := range bs.List {
		if b.Url == url {
			return b, nil
		}
	}

	return nil, nil
}

func Delete(bs *bookmark.Bookmarks, idx int) error {
	bs.List = slices.Delete(bs.List, idx, idx + 1)
	var data [][]string

	data = append(data, bookmark.Tags())
	data = append(data, bs.ToData()...)

	f, err := os.OpenFile(Path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)

	if err != nil {
		return err
	}

	defer f.Close()

	w := csv.NewWriter(f)

	if err := w.WriteAll(data); err != nil {
		return err
	}

	return nil
}

