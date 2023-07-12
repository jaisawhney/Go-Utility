package main

import (
	"strings"
	"testing"
)

func TestScrape(t *testing.T) {
	var tests = []struct {
		url  string
		html string
	}{
		{"https://example.com", "<h1>Example Domain</h1>"},
		{"https://news.ycombinator.com", "<a href=\"news\">Hacker News</a>"},
	}

	for _, test := range tests {
		if output := scrapeHtml(test.url); !strings.Contains(output, test.html) {
			t.Error("expected: {}\n\nrecieved: {}", test.html, output)
		}
	}
}

func TestCreateMarkdown(t *testing.T) {
	var tests = []struct {
		url      string
		markdown string
	}{
		{"https://example.com", "# Example Domain"},
		{"https://news.ycombinator.com", "**[Hacker News]"},
	}

	for _, test := range tests {
		html := scrapeHtml(test.url)
		if output := createMarkdown(test.url, html); !strings.Contains(output, test.markdown) {
			t.Error("expected: {}\n\nrecieved: {}", test.markdown, output)
		}
	}
}

func BenchmarkMarkdownScrape(b *testing.B) {
	// The time to scrape and convert to markdown

	html := scrapeHtml("https://example.com")
	createMarkdown("example.com", html)
}
