package bookmark

import (
	"fmt"
	"testing"
	"time"

	"bou.ke/monkey"
	"github.com/stretchr/testify/assert"
)


func TestTag(t *testing.T) {
	expect := []string{
		"datetime",
		"category",
		"title",
		"url",
	}

	result := Tags()
	assert.Equal(t, result, expect)
}

func TestCheckCategory(t *testing.T) {
	b := createTestBookmark()

	tests := []struct {
		input  string
		expect bool
	}{
		{"golang", true},
		{"python", false},
		{"Golang", false},
	}

	for _, tt := range tests {
		result := b.CheckCategory(tt.input)
		assert.Equal(t, result, tt.expect)
	}
}

func TestCheckKeyword(t *testing.T) {
	b := createTestBookmark()

	tests := []struct {
		input  string
		expect bool
	}{
		{"CLI", true},
		{"Cobra", true},
		{"CLIs", false},
		{"tool", false},
	}

	for _, tt := range tests {
		result := b.CheckKeyword(tt.input)
		assert.Equal(t, result, tt.expect)
	}
}

func TestBookmarkToData(t *testing.T) {
	expect := []string {
		"2025-05-30 01:35:43",
		"golang",
		"Creating a CLI in Go Using Cobra",
		"https://test.com/cde21ht",
	}

	b := createTestBookmark()
	result := b.ToData()
	assert.Equal(t, result, expect)
}

func TestBookmarksToData(t *testing.T) {
	expect := [][]string {
		{
			"2025-05-30 01:35:43",
			"golang",
			"Creating a CLI in Go Using Cobra",
			"https://test.com/cde21ht",
		},
		{
			"2025-05-31 17:35:43",
			"SQL",
			"Effective SQL",
			"https://test.com/avb27st",
		},
		{
			"2025-05-31 17:35:43",
			"python",
			"Effective Python",
			"https://test.com/vbwb27lq",
		},
	}

	bs := createTestBookmarks()
	result := bs.ToData()
	assert.Equal(t, result, expect)
}

func TestGetElement(t *testing.T) {
	bs := createTestBookmarks()

	tests := []struct {
		idx    int
		expectRet *Bookmark
		expectErr    error
	}{
		{0, &Bookmark{"2025-05-30 01:35:43", "golang", "Creating a CLI in Go Using Cobra", "https://test.com/cde21ht",}, nil},
		{1, &Bookmark{"2025-05-31 17:35:43", "SQL", "Effective SQL", "https://test.com/avb27st"}, nil},
		{-1, nil, fmt.Errorf("index out of range; the number of bookmarks: %d", len(bs.List))},
		{3, nil, fmt.Errorf("index out of range; the number of bookmarks: %d", len(bs.List))},
		{5, nil, fmt.Errorf("index out of range; the number of bookmarks: %d", len(bs.List))},
	}

	for _, tt := range tests {
		b, err := bs.GetElement(tt.idx)

		if tt.expectRet != nil {
			assert.Equal(t, tt.expectRet, b)
		} else {
			if assert.Error(t, err) {
				assert.Equal(t, tt.expectErr, err)
			}
		}
	}
}

func createTestBookmark() *Bookmark {
	b := &Bookmark{
		Datetime: "2025-05-30 01:35:43",
		Category: "golang",
		Title:    "Creating a CLI in Go Using Cobra",
		Url:      "https://test.com/cde21ht",
	}
	return b
}

func createTestBookmarks() *Bookmarks {
	li := []*Bookmark {
		{
			"2025-05-30 01:35:43",
			"golang",
			"Creating a CLI in Go Using Cobra",
			"https://test.com/cde21ht",
		},
		{
			"2025-05-31 17:35:43",
			"SQL",
			"Effective SQL",
			"https://test.com/avb27st",
		},
		{
			"2025-05-31 17:35:43",
			"python",
			"Effective Python",
			"https://test.com/vbwb27lq",
		},
	}
	bs := &Bookmarks {
		List: li,
	}
	return bs
}

func TestNew(t *testing.T) {
	t.Run("time mock test using monkey", func(t *testing.T) {
		mockTime := time.Date(2025, 6, 1, 1, 0, 0, 0, time.Local)
		patch := monkey.Patch(time.Now, func() time.Time {return mockTime})

		defer patch.Unpatch()
		
		expect := &Bookmark{"2025-06-01 01:00:00", "go", "go test", "https://test.com"}
		result := New("go", "go test", "https://test.com")
		assert.Equal(t, expect, result)
	
	})
}