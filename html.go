package autocorrect

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html"
)

// FormatHTML format HTML content
func FormatHTML(body string) (out string, err error) {
	return FormatHTMLWithOption(body, defaultOption)
}

// FormatHTMLWithOption format HTML content
func FormatHTMLWithOption(body string, opt Option) (out string, err error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(body))
	if err != nil {
		return body, err
	}

	traverseTextNodes(doc.First().Nodes[0], func(node *html.Node) {
		node.Data = FormatWithOption(node.Data, opt)
	})

	body, err = doc.Find("body").Html()
	if err != nil {
		return body, err
	}

	return body, nil
}

func traverseTextNodes(node *html.Node, fn func(*html.Node)) {
	if node == nil {
		return
	}
	if node.Type == html.TextNode {
		fn(node)
	}

	cur := node.FirstChild

	for cur != nil {
		next := cur.NextSibling
		traverseTextNodes(cur, fn)
		cur = next
	}
}
