package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	svg "github.com/ajstarks/svgo"
)

// API Info and formats
const (
	NYTAPIkey = "7f35400618208b235360ee874ed70df0"
	NYTfmt    = "http://api.nytimes.com/svc/news/v3/content/all/%s/"
	param     = ".json?api-key=%s&limit=5"
	style     = "font-size:20pt;font-family:sans-serif;fill:white"
	errfmt    = "unable to get network data for %s (%s)"
	datefmt   = "Monday Jan 2, 2006"
	usage     = "(arts, health, sports, science, technology, u.s., world)"
)

// NYTHeadlines is the headline info from the New York Times
type NYTHeadlines struct {
	Status     string   `json:"status"`
	Copyright  string   `json:"copyright"`
	NumResults int      `json:"num_results"`
	Results    []result `json:"results"`
}

type result struct {
	Section    string `json:"section"`
	Subsection string `json:"subsection"`
	Title      string `json:"title"`
	Abstract   string `json:"abstract"`
	Thumbnail  string `json:"thumbnail_standard"`
}

func main() {
	var section = flag.String("h", "u.s.", usage)
	flag.Parse()
	canvas := svg.New(os.Stdout)
	canvas.Start(1200, 900)
	canvas.Rect(0, 0, 1200, 900)
	nytheadlines(canvas, *section)
	canvas.End()
}

// netread derefernces a URL,
// returning the Reader, with an error
func netread(url string) (io.ReadCloser, error) {
	client := &http.Client{Timeout: 30 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(errfmt, url, resp.Status)
	}
	return resp.Body, nil
}

// nytheadlines retrieves data
// from the New York Times API, decodes and displays it.
func nytheadlines(canvas *svg.SVG, section string) {
	r, err := netread(fmt.Sprintf(NYTfmt+param, section, NYTAPIkey))
	if err != nil {
		fmt.Fprintf(os.Stderr, "headline read error: %v\n", err)
		return
	}
	defer r.Close()
	var data NYTHeadlines
	if err = json.NewDecoder(r).Decode(&data); err != nil {
		fmt.Fprintf(os.Stderr, "decode: %v\n", err)
		return
	}
	thumbsize := 75
	top, left := 200, 150
	imx, titley := left-100, top-100
	x, y := left, top
	ts := fmt.Sprintf("New York Times for %s", time.Now().Format(datefmt))
	canvas.Gstyle(style)
	canvas.Text(imx, titley, ts, "font-size:175%")
	for _, d := range data.Results {
		canvas.Text(x, y, d.Title)
		canvas.Image(imx, y-thumbsize/2, thumbsize, thumbsize, d.Thumbnail)
		y += 120
	}
	canvas.Gend()
}
