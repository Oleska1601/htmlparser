package usecase

import (
	"fmt"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

func nodeParser(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("http.Get: %w", err)
	}
	defer resp.Body.Close()
	doc, err := html.Parse(resp.Body)
	if err != nil {
		return "", fmt.Errorf("html.Parse: %w", err)
	}
	var buf strings.Builder
	var processNode func(*html.Node)
	processNode = func(n *html.Node) {
		if n.Type == html.ElementNode && (n.Data == "script" || n.Data == "style") {
			return
		}
		if n.Type == html.TextNode {
			text := strings.TrimSpace(n.Data)
			if text != "" {
				buf.WriteString(text)
				buf.WriteString("\n")
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			processNode(c)
		}
	}
	processNode(doc)
	return buf.String(), nil
}
