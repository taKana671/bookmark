package web

import (
	"fmt"
	"net/http"
	"os/exec"
	"runtime"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/spf13/cobra"
)

func GetTitle(url string) (string, error) {
	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	resp, err := client.Get(url)

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return "", fmt.Errorf("HTTP statue code: %s", resp.Status)

	}

	doc, _ := goquery.NewDocumentFromReader(resp.Body)
	title := doc.Find("title").Text()
	return title, nil

}

type openCommand struct {
	cmd     string
	args    []string
}


func newOpenCommand(nm string, url string) (*openCommand, error) {
	openCmd := openCommand{}

	switch nm {
		case "windows":
			openCmd.cmd = "rundll32.exe"
			openCmd.args = append(openCmd.args, "url.dll,FileProtocolHandler")

		case "linux":
			openCmd.cmd = "xdg-open"

		case "darwin":
			openCmd.cmd = "open"
		
		default:
			return &openCmd, fmt.Errorf("not supported : %s", nm)
		}

	openCmd.args = append(openCmd.args, url)
	return &openCmd, nil
}


func Open(cmd *cobra.Command, url string) error {
	openCmd, err := newOpenCommand(runtime.GOOS, url)

	if err != nil {
		return err
	}

	err = exec.Command(openCmd.cmd, openCmd.args...).Start()	
	// err := exec.Command("rundll32.exe", "url.dll,FileProtocolHandler", bm.Url).Start()
	
	if err != nil {
		return err
	}

	return nil
}