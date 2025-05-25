package web

import (
	"fmt"
	"net/http"
	"os/exec"
	"runtime"

	"github.com/PuerkitoBio/goquery"
	"github.com/spf13/cobra"
)

func GetTitle(cmd *cobra.Command, url string) (string, error) {
	var title string
	resp, err := http.Get(url)

	if err != nil {
		cmd.PrintErrf("failed to get title: %s, %s", url, err)
		return title, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		cmd.PrintErrf("failed to fetch data: %d %s", resp.StatusCode, resp.Status)
		return title, fmt.Errorf("HTTP statue code: %s", resp.Status)

	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)

	if err != nil {
		cmd.PrintErrf("failed to load html: %s", err)
		return title, err
	}

	title = doc.Find("title").Text()
	cmd.Println(title)
	return title, nil

}

func Open(cmd *cobra.Command, url string) error {
	var openCmd string
	var args []string

	switch runtime.GOOS {
		case "windows":
			openCmd = "rundll32.exe"
			args = append(args, "url.dll,FileProtocolHandler")
		case "linux":
			openCmd = "xdg-open"
		case "darwin":
			openCmd = "open"
		default:
			return fmt.Errorf("not supported : %s", runtime.GOOS)
		}

	args = append(args, url)
	err := exec.Command(openCmd, args...).Start()	
	// err := exec.Command("rundll32.exe", "url.dll,FileProtocolHandler", bm.Url).Start()
	
	if err != nil {
		cmd.PrintErrln("cannot open site")
		return err
	}

	return nil
}