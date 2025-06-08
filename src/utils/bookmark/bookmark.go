package bookmark

import (
	"fmt"
	"reflect"
	"strings"
	"time"
)

const layout = "2006-01-02 15:04:05"

type DtKey string
const DatetimeKey DtKey = "now"


type Bookmark struct {
	Datetime     string `csv:"datetime"`
	Category string `csv:"category"`
	Title    string `csv:"title"`
	Url      string `csv:"url"`
}

type Bookmarks struct {
	List []*Bookmark
}


func New(category, title, url string) *Bookmark {
	dt := time.Now()

	bm := &Bookmark{
		Datetime: dt.Format(layout),
		Category: category,
		Title: title,
		Url: url,
	}
	return bm
}

func Tags() []string {
	t := reflect.TypeOf(Bookmark{})
	tags := make([]string, t.NumField())
	
	for i := range t.NumField() {
		field := t.Field(i)
		tags[i] = field.Tag.Get("csv")
	}

	return tags
}

func (b *Bookmark) CheckCategory(c string) bool {
	if len(c) == 0 {
		return true
	}

	if b.Category == c {
		return true
	}
	return false
}

func (b *Bookmark) CheckKeyword(k string) bool {
	if len(k) == 0 {
		return true
	}

	if strings.Contains(b.Title, k) {
		return true
	}
	return false
}

func (b *Bookmark) ToData() []string {
	return []string{b.Datetime, b.Category, b.Title, b.Url}
}

func (bs *Bookmarks) ToData() [][]string {
	n := len(bs.List)
	data := make([][]string, n)

	for i := range n {
		b := bs.List[i]
		data[i] = b.ToData()
	}

	return data
}

func (bs *Bookmarks) GetElement(idx int) (*Bookmark, error) {
	if idx < 0 || idx >= len(bs.List) {
		return nil, fmt.Errorf("index out of range; the number of bookmarks: %d", len(bs.List))
	}

	return bs.List[idx], nil
}