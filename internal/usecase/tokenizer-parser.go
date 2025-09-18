package usecase

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

func tokenizerParser(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("http.Get: %w", err)
	}
	defer resp.Body.Close()
	var buf strings.Builder
	var skip bool
	tokenizer := html.NewTokenizer(resp.Body)
	for {
		tt := tokenizer.Next() // scans next token and return its type
		switch tt {
		case html.ErrorToken:
			if errors.Is(tokenizer.Err(), io.EOF) {
				return buf.String(), nil
			}
			return buf.String(), errors.New("html.ErrorToken")
		case html.StartTagToken:
			t := tokenizer.Token() // get current token
			if t.Data == "script" || t.Data == "style" {
				skip = true
			}
		case html.EndTagToken:
			t := tokenizer.Token() // get current token
			if t.Data == "script" || t.Data == "style" {
				skip = false
			}

		case html.TextToken:
			if !skip {
				text := strings.TrimSpace(string(tokenizer.Text()))
				if text != "" {
					buf.WriteString(text)
					buf.WriteString("\n")
				}
			}

		}
	}
}
