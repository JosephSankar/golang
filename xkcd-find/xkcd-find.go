package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type xkcd struct {
	Month      string `json:"month"`
	Num        int    `json:"num"`
	Year       string `json:"year"`
	Title      string `json:"title"`
	Transcript string `json:"transcript"`
	Day        string `json:"day"`
}

const NUM_COMICS = 2986

func main() {
	if len(os.Args) <= 2 {
		fmt.Println("A json file and search term(s) must be provided as command-line arguments")
		os.Exit(-1)
	}

	f, err := os.Open(os.Args[1])

	if err != nil {
		fmt.Printf("Could not open file %s due to error: %s", os.Args[1], err.Error())
		os.Exit(-1)
	}

	defer f.Close()

	reader := bufio.NewReader(f)
	xkcds := make([]xkcd, 0, NUM_COMICS)

	if err = json.NewDecoder(reader).Decode(&xkcds); err != nil {
		fmt.Printf("Could not decode file %s due to error: %s", os.Args[1], err.Error())
		os.Exit(-1)
	}

	for _, xkcd := range xkcds {
		termsInTitle := true
		termsInTranscript := true
		for _, term := range os.Args[2:] {
			termInTitle := false
			termInTranscript := false

			titleScan := bufio.NewScanner(strings.NewReader(xkcd.Title))
			titleScan.Split(bufio.ScanWords)
			transcriptScan := bufio.NewScanner(strings.NewReader(xkcd.Transcript))
			transcriptScan.Split(bufio.ScanWords)

			for titleScan.Scan() {
				if strings.EqualFold(strings.Trim(titleScan.Text(), "[].!,"), term) {
					termInTitle = true
					break
				}
			}

			for transcriptScan.Scan() {
				if strings.EqualFold(strings.Trim(transcriptScan.Text(), "[].!,"), term) {
					termInTranscript = true
					break
				}
			}

			if !termInTitle {
				termsInTitle = false
			}

			if !termInTranscript {
				termsInTranscript = false
			}

			if !termInTitle && !termInTranscript {
				break
			}
		}

		if termsInTitle || termsInTranscript {
			fmt.Printf("https://xkcd.com/%d/ %s/%s/%s \"%s\"\n", xkcd.Num, xkcd.Month, xkcd.Day, xkcd.Year, xkcd.Title)
		}
	}
}
