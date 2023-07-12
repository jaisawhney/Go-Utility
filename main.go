package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/chromedp/chromedp"

	md "github.com/JohannesKaufmann/html-to-markdown"
	"github.com/JohannesKaufmann/html-to-markdown/plugin"
)

// Navigates to and fetches the HTML for the provided website
func scrapeHtml(url string) string {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	var html string
	err := chromedp.Run(ctx,
		chromedp.Navigate(url),
		chromedp.OuterHTML("body", &html, chromedp.ByQuery),
	)
	if err != nil {
		panic(err)
	}
	return html
}

// Processes the HTML markup and converts it to a markdown string
func createMarkdown(baseDomain string, html string) string {
	converter := md.NewConverter(baseDomain, true, nil)
	converter.Use(plugin.Table())
	converter.Use(plugin.TaskListItems())

	markdown, err := converter.ConvertString(html)
	if err != nil {
		panic(err)
	}
	return markdown
}

// Saves the new markdown file
func saveFile(baseDomain string, markdown string, filePath string) {
	outputFile, err := os.Create("./output/" + baseDomain + ".md")
	if err != nil {
		panic(err)
	}

	outputFile.Write([]byte(markdown))
	fmt.Println("Saved to '" + baseDomain + ".md'")
}

func main() {
	url := flag.String("url", "", "The URL of the website to download")
	flag.Parse()

	if *url != "" {
		var html = scrapeHtml(*url)
		var baseDomain = md.DomainFromURL(*url)
		var markdown = createMarkdown(baseDomain, html)

		saveFile(baseDomain, markdown, markdown)
	} else {
		fmt.Println("'--url' flag missing")
	}
}
