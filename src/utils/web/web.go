package web

import (
	"fmt"
	"net/http"
	"os/exec"
	"time"

	"github.com/PuerkitoBio/goquery"
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

type OpenCommand struct {
	Cmd     string
	Args    []string
}

func NewOpenCommand(nm string, url string) (*OpenCommand, error) {
	openCmd := OpenCommand{}

	switch nm {
		case "windows":
			openCmd.Cmd = "rundll32.exe"
			openCmd.Args = append(openCmd.Args, "url.dll,FileProtocolHandler")

		case "linux":
			openCmd.Cmd = "xdg-open"

		case "darwin":
			openCmd.Cmd = "open"
		
		default:
			return &openCmd, fmt.Errorf("not supported : %s", nm)
		}

	openCmd.Args = append(openCmd.Args, url)
	return &openCmd, nil
}


func (o *OpenCommand) Execute() error {
	if err := exec.Command(o.Cmd, o.Args...).Start(); err != nil {
		return err
	}
	return nil
}