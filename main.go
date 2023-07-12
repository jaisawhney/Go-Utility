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
func fetchHtml(website string) string {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	var html string
	chromedp.Run(ctx,
		chromedp.Navigate(website),
		chromedp.OuterHTML("body", &html, chromedp.ByQuery),
	)
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
	website := flag.String("website", "", "Website to download")

	flag.Parse()

	if *website != "" {
		var html = fetchHtml(*website)
		var baseDomain = md.DomainFromURL(*website)
		var markdown = createMarkdown(baseDomain, html)

		saveFile(baseDomain, markdown, markdown)
	} else {
		fmt.Println("'--website' flag missing")
	}
}
