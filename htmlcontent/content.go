package htmlcontent

import (
	"regexp"
	"strings"

	"golang.org/x/net/html"
)

func GetTitleFromHTML(htmlContent string) string {
	reader := strings.NewReader(htmlContent)
	tokenizer := html.NewTokenizer(reader)

	for {
		tokenType := tokenizer.Next()

		switch tokenType {
		case html.ErrorToken:
			return "" // End of the document
		case html.StartTagToken, html.SelfClosingTagToken:
			token := tokenizer.Token()

			if token.Data == "title" {
				// Found the start of title tag, now extract the text content
				return extractTextContent(tokenizer)
			}
		}
	}
}

func extractTextContent(tokenizer *html.Tokenizer) string {
	var textContent string

	for {
		tokenType := tokenizer.Next()

		switch tokenType {
		case html.ErrorToken, html.EndTagToken:
			return textContent // End of title tag or document
		case html.TextToken:
			token := tokenizer.Token()
			textContent += token.Data
		}
	}
}

func GetHTMLBetweenMarkers(htmlContent string) string {
	startPattern := `<div id="topic_content" class="topic-content markdown-body">`
	endPattern := `<span id="mark-text">点击收藏 </span>`

	re := regexp.MustCompile(startPattern + `([\s\S]*?)` + endPattern)
	match := re.FindStringSubmatch(htmlContent)

	if len(match) >= 2 {
		extractedHTML := match[1]
		return extractedHTML
	}

	return ""
}
