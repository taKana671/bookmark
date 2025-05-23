package csv_handler

import (
	"bookmark/src/utils/bookmark"
	"bytes"
	"encoding/csv"
	"log"
	"os"

	"github.com/jszwec/csvutil"
)

const PATH = "bookmarks.csv"


func Write(data []string) error {
	f, err := os.OpenFile(PATH, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		log.Printf("failed to open %s: %s", PATH, err)
		return err
	}

	defer f.Close()

	w := csv.NewWriter(f)
	if err := w.Write(data); err != nil {
		log.Printf("failed to write %s: %s", PATH, err)
		return err
	}

	w.Flush()
	return nil
}

func Read() ([]*bookmark.Bookmark, error) {
	var bms []*bookmark.Bookmark

	f, err := os.Open(PATH)
	
	if err != nil {
		log.Fatal(err)
		return bms, err
	}

	defer f.Close()
	
	r := csv.NewReader(f)
	rows, err := r.ReadAll()

	if err != nil {
		log.Fatal(err)
		return bms, err
	}

	b := &bytes.Buffer{}
	if err := csv.NewWriter(b).WriteAll(rows); err != nil {
		log.Fatal(err)
		return bms, err
	}

	if err := csvutil.Unmarshal(b.Bytes(), &bms); err != nil {
		log.Fatal(err)
		return bms, err
	}

	return bms, nil

}