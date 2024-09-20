package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

type xkcd struct {
	Month      string `json:"month"`
	Num        int    `json:"num"`
	Year       string `json:"year"`
	Title      string `json:"title"`
	Transcript string `json:"transcript"`
	Day        string `json:"day"`
}

const NUM_COMICS = 2987

func main() {
	if len(os.Args) == 1 {
		fmt.Println("An output filename must be specified as a command-line parameter")
		os.Exit(-1)
	}

	const base = "https://xkcd.com/"
	const suffix = "/info.0.json"

	xkcds := make([]xkcd, 0, NUM_COMICS)

	for i := 1; i <= NUM_COMICS; i++ {
		url := base + strconv.Itoa(i) + suffix

		resp, err := http.Get(url)

		if err != nil {
			fmt.Printf("skipping %d: got error %s\n", i, err.Error())
			continue
		}

		if resp.StatusCode != http.StatusOK {
			fmt.Printf("skipping %d: got %d\n", i, resp.StatusCode)
			continue
		}

		defer resp.Body.Close()

		var item xkcd

		if err = json.NewDecoder(resp.Body).Decode(&item); err != nil {
			fmt.Printf("skipping %d because we could not decode it", i)
			continue
		}

		xkcds = append(xkcds, item)
	}

	f, err := os.Create(os.Args[1])

	if err != nil {
		fmt.Println("Could not open file " + os.Args[1] + " for writing")
		return
	}

	defer f.Close()

	w := bufio.NewWriter(f)

	enc := json.NewEncoder(w)

	enc.SetIndent("", "    ")

	if err = enc.Encode(xkcds); err != nil {
		fmt.Println("error encoding list of xkcds")
		return
	}

	w.Flush()

	fmt.Printf("Read %d comics", len(xkcds))
}
