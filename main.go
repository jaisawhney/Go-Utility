package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/chromedp/cdproto/dom"
	"github.com/chromedp/chromedp"
)

func main() {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	website := flag.String("website", "", "Website to download")

	flag.Parse()

	if *website != "" {
		var html string
		chromedp.Run(ctx,
			chromedp.Navigate(*website),
			chromedp.ActionFunc(func(ctx context.Context) error {
				node, err := dom.GetDocument().Do(ctx)
				if err != nil {
					return err
				}
				html, err = dom.GetOuterHTML().WithNodeID(node.NodeID).Do(ctx)
				return err
			}),
		)
		fmt.Println(html)
	} else {
		fmt.Println("'--website' flag missing")
	}
}
