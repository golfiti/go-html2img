// How to run
// go run html2img.go report.html export.png

package main

import (
	"context"
	"io/ioutil"
	"log"
	"os"

	"github.com/chromedp/chromedp"
)

func main() {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	var err = chromedp.Run(ctx, chromedp.EmulateViewport(1024, 768, chromedp.EmulateScale(2.0)))
	if err != nil {
		log.Fatal(err)
	}

	inputPath := os.Args[1]
	if err := chromedp.Run(ctx, chromedp.Navigate(inputPath)); err != nil {
		log.Fatal(err)
	}

	var buf []byte
	if err := chromedp.Run(ctx, chromedp.CaptureScreenshot(&buf)); err != nil {
		log.Fatal(err)
	}

	outputPath := os.Args[2]
	if err := ioutil.WriteFile(outputPath, buf, os.ModePerm); err != nil {
		log.Fatal(err)
	}
}
