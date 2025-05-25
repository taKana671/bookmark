package bookmark

import (
	"strings"
	"time"
)

const layout = "2006-01-02 15:04:05"

type Bookmark struct {
	Date     string `csv:"date"`
	Category string `csv:"category"`
	Title    string `csv:"title"`
	Url      string `csv:"url"`
}


func New(category, title, url string) *Bookmark {
	t := time.Now()

	bm := &Bookmark{
		Date: t.Format(layout),
		Category: category,
		Title: title,
		Url: url,
	}
	return bm
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
	return []string{b.Date, b.Category, b.Title, b.Url}
}

func (b *Bookmark) Fields() []string {
	return []string{"date", "category", "title", "url"}
	// how to get tag name, see below
	// https://text.baldanders.info/golang/struct-tag/
} 