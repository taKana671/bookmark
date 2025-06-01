package web

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)


func TestGetTitle(t *testing.T) {
	mockTimeout := httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			time.Sleep(10 * time.Second)
			w.WriteHeader(http.StatusOK)
		},
	))
	
	defer mockTimeout.Close()

	mockInternalServerError := httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusInternalServerError)
		},
	))

	defer mockInternalServerError.Close()

	mockStatusOK := httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w, "<html><body><title>test title</title></html>")
		},
	))

	defer mockStatusOK.Close()

	// timeout
	title, err := GetTitle(mockTimeout.URL)
	assert.Empty(t, title)
	assert.NotEmpty(t, err)
	
	// 500
	expectErr := fmt.Errorf("HTTP statue code: %s", "500 Internal Server Error")
	title, err = GetTitle(mockInternalServerError.URL)
	assert.Empty(t, title)
	assert.Equal(t, expectErr, err)

	// successfully get title
	title, err = GetTitle(mockStatusOK.URL)
	assert.Equal(t, title, "test title")
	assert.Empty(t, err)

}

func TestNewOpenCommand(t *testing.T) {
	tests := []struct {
		nm    string
		url   string
		expectCmd string
		expectArgs []string
	}{
		{"windows", "https://test", "rundll32.exe", []string{"url.dll,FileProtocolHandler", "https://test"}},
		{"linux", "https://test", "xdg-open", []string{"https://test"}},
		{"darwin", "https://test", "open", []string{"https://test"}},
		{"aaa", "https://test", "", nil},

	}

	for _, tt := range tests {
		openCmd, err := newOpenCommand(tt.nm, tt.url)

		if err == nil {
			assert.Equal(t, tt.expectCmd, openCmd.cmd)
		} else {
			if assert.Error(t, err) {
				assert.Equal(t, err, fmt.Errorf("not supported : %s", tt.nm))
			}
		}
	}
}


// var testCase string

// func fakeExecCommand(command string, args ...string) *exec.Cmd {
// 	cs := []string{"-test.run=TestHelperProcess", "--", command}
// 	cs = append(cs, args...)
// 	cmd := exec.Command(os.Args[0], cs...)
// 	tc := "TEST_CASE=" + testCase

// 	cmd.Env = []string{"GO_WANT_HELPER_PROCESS=1", tc}
//   	return cmd
// }

// func TestOpen(t *testing.T) {
// 	execCommand = fakeExecCommand
// 	defer func() {execCommand}
// }